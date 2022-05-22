package internal

import "github.com/rjeczalik/notify"

var (
	// EventTypePut contains the notify events that will cause a put (writer)
	EventTypePut = []notify.Event{notify.Create, notify.Write, notify.Rename, notify.FileNotifyChangeFileName, notify.FileNotifyChangeDirName}
	// EventTypeDelete contains the notify events that will cause a delete (remove)
	EventTypeDelete = []notify.Event{notify.Remove}
	// EventTypeGet contains the notify events that will cause a get (read)
	EventTypeGet = []notify.Event{notify.FileNotifyChangeLastAccess}
)

// IsGetEvent checks if the event return is a get event.
func IsGetEvent(event notify.Event) bool {
	return event&notify.FileNotifyChangeLastAccess != 0
}

// IsPutEvent checks if the event returned is a put event
func IsPutEvent(event notify.Event) bool {
	if event&notify.FileActionRenamedOldName != 0 {
		return false
	}
	for _, ev := range EventTypePut {
		if event&ev != 0 {
			return true
		}
	}
	return false
}

// IsDeleteEvent checks if the event returned is a delete event
func IsDeleteEvent(event notify.Event) bool {
	return event&notify.Remove != 0 || event&notify.FileActionRenamedOldName != 0
}

// getAllXattrs returns the extended attributes for a file if supported
// by the OS
func getAllXattrs(path string) (map[string]string, error) {
	return nil, nil
}
