package internal

import (
	"sync"
)

// S3Client construct
type S3Client struct {
	sync.Mutex
}
