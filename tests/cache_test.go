package tests

import (
	"testing"

	"github.com/gophers-latam/GoKey/gokey"
)

var operations gokey.Operations = new(gokey.Cache)

func TestCacheUpsert(t *testing.T) {
	_, err := operations.Upsert("key", []byte("value"), -1)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = operations.Upsert("", []byte("value"), 1)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCacheGet(t *testing.T) {
	_, err := operations.Upsert("key", []byte("value"), 10)
	if err != nil {
		t.Error("Expected no errors in Upsert method, got:", err)
	}

	value, err := operations.Get("key")
	if err != nil {
		t.Error("Expected no errors in Get method, got:", err)
	}

	if value != nil {
		t.Error("Expected a value, got nil")
	}
}

func TestCacheGetEmptyKey(t *testing.T) {
	_, err := operations.Get("")
	if err == nil {
		t.Error("Expected empty key error message, got nil")
	}
}

func TestCacheGetUnknowKey(t *testing.T) {
	_, err := operations.Get("Key")
	if err == nil {
		t.Error("Expected 'no related values' error message, got nil")

    // case ok
	_, err = operations.Upsert("key", []byte("value"), 0)
	if err != nil {
		t.Error("here: unexpected error")
	}
}