package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/gophers-latam/GoKey/gokey"
)

var operations gokey.Operations = new(gokey.Cache)

// go test -run TestCacheUpsert -v
func TestCacheUpsert(t *testing.T) {
	_, err := operations.Upsert("key", []byte("value"), -1)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = operations.Upsert("", []byte("value"), 1)
	if err == nil {
		t.Error(err.Error())
	}
}

// go test -run TestCacheGet -v
func TestCacheGet(t *testing.T) {
	_, err := operations.Upsert("key", []byte("value"), 10*time.Second)
	if err != nil {
		t.Error("Expected no errors in Upsert method, got:", err.Error())
	}

	value, err := operations.Get("key")
	if err != nil {
		t.Error("Expected no errors in Get method, got:", err.Error())
	}

	if value == nil {
		t.Error("Expected a value, got nil")
	}
}

// go test -run TestCacheGetExpiredKey -v
func TestCacheGetExpiredKey(t *testing.T) {
	_, err := operations.Upsert("key", []byte("value"), 1*time.Second)
	if err != nil {
		t.Error("expected no errors in Upsert method, got:", err.Error())
	}
	time.Sleep(1 * time.Second)

	_, err = operations.Get("key")
	if err == nil {
		t.Error("expected ErrExpiredKey, got: nil")
	}
	if err != nil {
		if !errors.Is(err, gokey.ErrExpiredKey) {
			t.Error("expected ErrExpiredKey, got:", err.Error())
		}
	}

	_, err = operations.Get("key")
	if err != nil {
		if !errors.Is(err, gokey.ErrNoExistKey) {
			t.Error("expected 'key does not exist' error message, got:", err.Error())
		}
	}
}

// go test -run TestCacheGetEmptyKey -v
func TestCacheGetEmptyKey(t *testing.T) {
	_, err := operations.Get("")
	if err == nil {
		t.Error("Expected empty key error message, got nil")
	}
}

// go test -run TestCacheGetUnknowKey -v
func TestCacheGetUnknowKey(t *testing.T) {
	_, err := operations.Get("Key")
	if err == nil {
		t.Error("Expected 'no related values' error message, got nil")
	}

	// case ok
	_, err = operations.Upsert("key", []byte("value"), 0)
	if err != nil {
		t.Error("here: unexpected error")
	}
}
