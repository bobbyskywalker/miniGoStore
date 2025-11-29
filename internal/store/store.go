package store

import (
	"sync"
	"time"
)

type Store struct {
	mutex sync.RWMutex
	data  map[string]ValueEntry
}

type ValueEntry struct {
	Value     []byte
	ExpiresAt time.Time
}
