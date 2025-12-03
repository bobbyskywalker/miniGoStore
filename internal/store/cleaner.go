package store

import (
	"log"
	"time"
)

const NumKeysToClean = 10

func (s *Store) StartCleaner() {
	go func() {
		log.Println("Cleaner routine started")
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			s.mutex.Lock()

			for i := 0; i < NumKeysToClean; i++ {
				if len(s.ttlKeys) == 0 {
					break
				}

				for k := range s.ttlKeys {
					entry := s.data[k]

					if entry.HasExpiry && time.Now().After(entry.ExpiresAt) {
						delete(s.data, k)
						delete(s.ttlKeys, k)
						log.Printf("Cleaner: removed key: %s", k)
					}
					break
				}
			}
			s.mutex.Unlock()
		}
	}()
}
