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

func (s *Store) Set(
	key string,
	value []byte,
	setOnExistent bool,
	setOnNonExistent bool,
	ttl *time.Time,
	retrievePrevious bool,
) []byte {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	oldEntry, exists := s.data[key]
	var oldValue []byte
	if exists {
		oldValue = oldEntry.Value
	}

	if setOnNonExistent && exists {
		return []byte("(nil)")
	}

	if setOnExistent && !exists {
		return []byte("(nil)")
	}

	entry := ValueEntry{Value: value}
	if ttl != nil {
		entry.ExpiresAt = *ttl
	}
	s.data[key] = entry

	if retrievePrevious {
		if oldValue == nil {
			return []byte("(nil)")
		}
		return oldValue
	}

	return []byte("OK")
}
