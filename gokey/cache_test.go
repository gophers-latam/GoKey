package gokey

import (
	"testing"
	"time"
)

var operations Operations = &Cache{pairsSet: map[string]pair{}}

// go test -run TestCacheUpsert -v
func TestCacheUpsert(t *testing.T) {
	_, err := operations.Upsert("key", []byte("value"), -1)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = operations.Upsert("", []byte("value"), 1)
	if err == nil {
		t.Error("key cannot be empty")
	}
}

// go test -run TestCacheGet -v
func TestCacheGet(t *testing.T) {
	_, err := operations.Upsert("key", []byte("value"), 10*time.Second)
	if err != nil {
		t.Error("expected no errors in Upsert method, got:", err.Error())
	}

	value, err := operations.Get("key")
	if err != nil {
		t.Error("expected no errors in Get method, got:", err.Error())
	}

	if value == nil {
		t.Error("expected a value, got nil")
	}
}

// go test -run TestCacheGetEmptyKey -v
func TestCacheGetEmptyKey(t *testing.T) {
	_, err := operations.Get("")
	if err == nil {
		t.Error("expected empty key error message, got nil")
	}
}

// go test -run TestCacheGetUnknownKey -v
func TestCacheGetUnknownKey(t *testing.T) {
	_, err := operations.Get("Key")
	if err == nil {
		t.Error("expected 'no related values' error message, got nil")
	}

	// case ok
	_, err = operations.Upsert("key", []byte("value"), 0)
	if err != nil {
		t.Error("here: unexpected error")
	}
}

// go test -run TestUpsertSameKey -v
func TestUpsertSameKey(t *testing.T) {
	key := "key"
	value := "value"
	newValue := "newValue"
	_, err := operations.Upsert(key, []byte(value), -1)

	if err != nil {
		t.Error("err should be nil", err)
	}

	v, err := operations.Get(key)

	if err != nil {
		t.Error("err should be nil", err)
	}

	if string(v) != value {
		t.Errorf("got different value from cache. Expected: %s, got: %s", value, string(v))
	}

	_, err = operations.Upsert(key, []byte(newValue), -1)
	v, err = operations.Get(key)

	if string(v) != newValue {
		t.Errorf("got different value from cache. Expected: %s, got: %s", newValue, string(v))
	}
}
