package store

import (
	"log/slog"
	"miniGoStore/internal/replies"
	"sync"
	"time"
)

type Store struct {
	mutex   sync.RWMutex
	data    map[string]ValueEntry
	ttlKeys map[string]struct{}
}

type ValueEntry struct {
	Value     []byte
	ExpiresAt time.Time
	HasExpiry bool
}

func NewStore() *Store {
	return &Store{
		data:    make(map[string]ValueEntry),
		ttlKeys: make(map[string]struct{}),
	}
}

func (s *Store) Get(key string) (ValueEntry, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	defer slog.Info("Retrieved key", slog.String("key", key))

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
		return []byte(replies.SetFailReply)
	}

	if setOnExistent && !exists {
		return []byte(replies.SetFailReply)
	}

	entry := ValueEntry{Value: value}
	if ttl != nil {
		entry.ExpiresAt = *ttl
		entry.HasExpiry = true
		s.ttlKeys[key] = struct{}{}
	} else {
		entry.HasExpiry = false
		delete(s.ttlKeys, key)
	}
	s.data[key] = entry

	if retrievePrevious {
		if oldValue == nil {
			return []byte(replies.SetFailReply)
		}
		return oldValue
	}

	slog.Debug("Set new key", slog.String("key", key))

	return []byte(replies.SuccessReply)
}

func (s *Store) SetEx(key string, ttl *time.Time, persist bool) (ValueEntry, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	defer slog.Debug("Retrieved key", slog.String("key", key))

	oldEntry, exists := s.data[key]

	if exists && ttl != nil {
		oldEntry.ExpiresAt = *ttl
		oldEntry.HasExpiry = true
		s.ttlKeys[key] = struct{}{}
		slog.Debug("Set new expiration time for key", slog.String("key", key))
	}
	if exists && persist {
		oldEntry.HasExpiry = false
		delete(s.ttlKeys, key)
		slog.Debug("Removed expiration date for key", slog.String("key", key))
	}
	s.data[key] = oldEntry

	return oldEntry, exists
}

func (s *Store) Exists(key string) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	defer slog.Debug("Checked for existence", slog.String("key", key))

	_, ok := s.data[key]
	if ok {
		return 1
	}
	return 0
}

func (s *Store) Del(keys []string) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	totalDeleted := 0

	for _, key := range keys {
		_, ok := s.data[key]
		if ok {
			delete(s.data, key)
			delete(s.ttlKeys, key)
			totalDeleted++
		}
	}

	slog.Debug("Deleted keys", slog.Int("total", totalDeleted))

	return totalDeleted
}

func (s *Store) CheckTtl(key string) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	defer slog.Debug("Retrieved TTL for key", slog.String("key", key))

	v, ok := s.data[key]
	if !ok {
		return -2
	}
	if !v.HasExpiry {
		return -1
	}

	return v.ExpiresAt.Second() - time.Now().Second()
}
