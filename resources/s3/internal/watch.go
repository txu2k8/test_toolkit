package internal

import (
	"context"
	"sync"
	"time"

	"github.com/minio/mc/pkg/probe"
	"github.com/minio/minio-go/v7/pkg/notification"
)

// EventInfo contains the information of the event that occurred and the source
// IP:PORT of the client which triggerred the event.
type EventInfo struct {
	Time         string
	Size         int64
	UserMetadata map[string]string
	Path         string
	Host         string
	Port         string
	UserAgent    string
	Type         notification.EventType
}

// WatchOptions contains watch configuration options
type WatchOptions struct {
	Prefix    string
	Suffix    string
	Events    []string
	Recursive bool
}

// WatchObject captures watch channels to read and listen on.
type WatchObject struct {
	// eventInfo will be put on this chan
	EventInfoChan chan []EventInfo
	// errors will be put on this chan
	ErrorChan chan *probe.Error
	// will stop the watcher goroutines
	DoneChan chan struct{}
}

// Events returns the chan receiving events
func (w *WatchObject) Events() chan []EventInfo {
	return w.EventInfoChan
}

// Errors returns the chan receiving errors
func (w *WatchObject) Errors() chan *probe.Error {
	return w.ErrorChan
}

// Watcher can be used to have one or multiple clients watch for notifications
type Watcher struct {
	sessionStartTime time.Time

	// all error will be added to this chan
	ErrorChan chan *probe.Error
	// all events will be added to this chan
	EventInfoChan chan []EventInfo

	// array of watchers joined
	o []*WatchObject

	// all watchers joining will enter this waitgroup
	wg sync.WaitGroup
}

// NewWatcher creates a new watcher
func NewWatcher(sessionStartTime time.Time) *Watcher {
	return &Watcher{
		sessionStartTime: sessionStartTime,
		ErrorChan:        make(chan *probe.Error),
		EventInfoChan:    make(chan []EventInfo),
		o:                []*WatchObject{},
	}
}

// Errors returns a channel which will receive errors
func (w *Watcher) Errors() chan *probe.Error {
	return w.ErrorChan
}

// Events returns a channel which will receive events
func (w *Watcher) Events() chan []EventInfo {
	return w.EventInfoChan
}

// Stop all watchers
func (w *Watcher) Stop() {
	for _, w := range w.o {
		close(w.DoneChan)
	}
	w.wg.Wait()
}

// Join the watcher with client
func (w *Watcher) Join(ctx context.Context, client Client, recursive bool) *probe.Error {
	wo, err := client.Watch(ctx, WatchOptions{
		Recursive: recursive,
		Events:    []string{"put", "delete", "bucket-creation", "bucket-removal"},
	})
	if err != nil {
		return err
	}

	w.o = append(w.o, wo)

	// join monitoring waitgroup
	w.wg.Add(1)

	// wait for events and errors of individual client watchers
	// and sent then to eventsChan and errorsChan
	go func() {
		defer w.wg.Done()

		for {
			select {
			case <-wo.DoneChan:
				return
			case events, ok := <-wo.Events():
				if !ok {
					return
				}
				w.EventInfoChan <- events
			case err, ok := <-wo.Errors():
				if !ok {
					return
				}

				w.ErrorChan <- err
			}
		}
	}()

	return nil
}
