package store

import (
	"bytes"
	"miniGoStore/internal/replies"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	store := NewStore()

	existentKey := "aa"
	store.data[existentKey] = ValueEntry{}

	nonExistentKey := "bb"

	_, ok := store.Get(existentKey)
	if !ok {
		t.Errorf("Incorrect result for existent key: not found in store")
	}

	_, ok = store.Get(nonExistentKey)
	if ok {
		t.Errorf("Incorrect result for non-existent key: found in store")
	}
}

func TestBasicSet(t *testing.T) {
	store := NewStore()

	rpl := store.Set("key", []byte("value"), false, false, nil, false)
	if !bytes.Equal(rpl, []byte(replies.SuccessReply)) {
		t.Errorf("Incorrect result: SET on empty store did not send OK reply")
	}

	_, ok := store.Get("key")
	if !ok {
		t.Errorf("Incorrect result: failed to GET a SET value")
	}
}

func TestSet_NX_OnExistingKey_FailsAndKeepsOldValue(t *testing.T) {
	store := NewStore()
	store.Set("key", []byte("old"), false, false, nil, false)

	// NX = true
	rpl := store.Set("key", []byte("new"), false, true, nil, false)
	if !bytes.Equal(rpl, []byte(replies.SetFailReply)) {
		t.Fatalf("expected FAIL reply for NX on existing key, got %q", rpl)
	}

	entry, ok := store.Get("key")
	if !ok {
		t.Fatalf("expected key to still exist")
	}
	if !bytes.Equal(entry.Value, []byte("old")) {
		t.Fatalf("expected value to remain %q, got %q", "old", entry.Value)
	}
}

func TestSet_NX_OnMissingKey_Succeeds(t *testing.T) {
	store := NewStore()

	// NX /w missing key
	rpl := store.Set("key", []byte("value"), false, true, nil, false)
	if !bytes.Equal(rpl, []byte(replies.SuccessReply)) {
		t.Fatalf("expected OK reply for NX on missing key, got %q", rpl)
	}

	entry, ok := store.Get("key")
	if !ok {
		t.Fatalf("expected key to exist")
	}
	if !bytes.Equal(entry.Value, []byte("value")) {
		t.Fatalf("expected value %q, got %q", "value", entry.Value)
	}
}

func TestSet_XX_OnMissingKey_Fails(t *testing.T) {
	store := NewStore()

	// XX = true
	rpl := store.Set("key", []byte("value"), true, false, nil, false)
	if !bytes.Equal(rpl, []byte(replies.SetFailReply)) {
		t.Fatalf("expected FAIL reply for XX on missing key, got %q", rpl)
	}

	if _, ok := store.Get("key"); ok {
		t.Fatalf("did not expect key to be created on XX with missing key")
	}
}

func TestSet_XX_OnExistingKey_SucceedsAndOverwrites(t *testing.T) {
	store := NewStore()
	store.Set("key", []byte("old"), false, false, nil, false)

	// XX with existing key
	rpl := store.Set("key", []byte("new"), true, false, nil, false)
	if !bytes.Equal(rpl, []byte(replies.SuccessReply)) {
		t.Fatalf("expected OK reply for XX on existing key, got %q", rpl)
	}

	entry, ok := store.Get("key")
	if !ok {
		t.Fatalf("expected key to exist")
	}
	if !bytes.Equal(entry.Value, []byte("new")) {
		t.Fatalf("expected value %q, got %q", "new", entry.Value)
	}
}

func TestSet_WithTTL_SetsExpiryAndMarksTTLKey(t *testing.T) {
	store := NewStore()
	exp := time.Now().Add(5 * time.Minute)

	rpl := store.Set("key", []byte("value"), false, false, &exp, false)
	if !bytes.Equal(rpl, []byte(replies.SuccessReply)) {
		t.Fatalf("expected OK reply, got %q", rpl)
	}

	entry, ok := store.Get("key")
	if !ok {
		t.Fatalf("expected key to exist")
	}
	if !entry.HasExpiry {
		t.Fatalf("expected HasExpiry to be true")
	}
	if !entry.ExpiresAt.Equal(exp) {
		t.Fatalf("expected ExpiresAt = %v, got %v", exp, entry.ExpiresAt)
	}

	if _, exists := store.ttlKeys["key"]; !exists {
		t.Fatalf("expected key to be present in ttlKeys")
	}
}

func TestSet_WithoutTTL_ClearsExistingExpiryAndTTLKey(t *testing.T) {
	store := NewStore()
	exp := time.Now().Add(5 * time.Minute)

	store.Set("key", []byte("old"), false, false, &exp, false)

	rpl := store.Set("key", []byte("new"), false, false, nil, false)
	if !bytes.Equal(rpl, []byte(replies.SuccessReply)) {
		t.Fatalf("expected OK reply, got %q", rpl)
	}

	entry, ok := store.Get("key")
	if !ok {
		t.Fatalf("expected key to exist")
	}

	if entry.HasExpiry {
		t.Fatalf("expected HasExpiry to be false after SET without TTL")
	}

	if _, exists := store.ttlKeys["key"]; exists {
		t.Fatalf("expected key to be removed from ttlKeys")
	}
}

func TestSet_RetrievePrevious_NoPrevious_ReturnsFailReply(t *testing.T) {
	store := NewStore()

	// retrievePrevious = true
	rpl := store.Set("key", []byte("value"), false, false, nil, true)
	if !bytes.Equal(rpl, []byte(replies.SetFailReply)) {
		t.Fatalf("expected FAIL reply when retrieving previous on non-existent key, got %q", rpl)
	}

	entry, ok := store.Get("key")
	if !ok {
		t.Fatalf("expected key to still be created")
	}
	if !bytes.Equal(entry.Value, []byte("value")) {
		t.Fatalf("expected value %q, got %q", "value", entry.Value)
	}
}

func TestSet_RetrievePrevious_WithPrevious_ReturnsOldValueAndStoresNew(t *testing.T) {
	store := NewStore()

	store.Set("key", []byte("old"), false, false, nil, false)

	rpl := store.Set("key", []byte("new"), false, false, nil, true)
	if !bytes.Equal(rpl, []byte("old")) {
		t.Fatalf("expected returned value to be old value %q, got %q", "old", rpl)
	}

	entry, ok := store.Get("key")
	if !ok {
		t.Fatalf("expected key to exist")
	}
	if !bytes.Equal(entry.Value, []byte("new")) {
		t.Fatalf("expected value %q after overwrite, got %q", "new", entry.Value)
	}
}

func TestExists_ReturnsCorrectFlag(t *testing.T) {
	store := NewStore()

	store.data["key"] = ValueEntry{}

	exists := store.Exists("key")
	if exists == 0 {
		t.Fatalf("exists should return 1 for existent key in store")
	}

	exists = store.Exists("key1")
	if exists == 1 {
		t.Fatalf("exists should return 0 for non-existent key in store")
	}
}

func TestDel_DeletesExistentKeys(t *testing.T) {
	store := NewStore()

	store.data["key"] = ValueEntry{}
	store.data["abc"] = ValueEntry{}

	keys := []string{"key", "abc", "ooo"}

	numDel := store.Del(keys)
	if numDel != 2 {
		t.Fatalf("del should delete two existent keys")
	}
}

func TestBasicSetEx(t *testing.T) {
	store := NewStore()

	ttl := time.Now().Add(10 * time.Second)
	store.data["key"] = ValueEntry{Value: []byte("val"), ExpiresAt: ttl, HasExpiry: true}

	newTtl := time.Now().Add(20 * time.Second)
	_, ok := store.SetEx("key", &newTtl, false)
	if !ok {
		t.Fatalf("setex should return the value")
	}

	if store.data["key"].ExpiresAt != newTtl {
		t.Fatalf("Setex should update expiration time")
	}
}

func TestSetExPersistShouldRmTtl(t *testing.T) {
	store := NewStore()

	ttl := time.Now().Add(10 * time.Second)
	store.data["key"] = ValueEntry{Value: []byte("val"), ExpiresAt: ttl, HasExpiry: true}

	_, ok := store.SetEx("key", nil, true)
	if !ok {
		t.Fatalf("setex should return the value")
	}

	if store.data["key"].HasExpiry {
		t.Fatalf("Setex should make hasexpiry false for key")
	}
}
