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

func NewStore() *Store {
	return &Store{
		data: make(map[string]ValueEntry),
	}
}

func (s *Store) Get(key string) (ValueEntry, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	v, ok := s.data[key]
	return v, ok
}

func (s *Store) Set(key string, value []byte) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// todo
}
