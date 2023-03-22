package gokey

import (
	"testing"
)

// go test -run TestTupleMaxSize -v -limit 50 -pairs 1000
func TestTupleMaxSize(t *testing.T) {
	_, err := operations.Upsert("key", []byte("more-than-32-byte-string-key-value-tuple"), -1)
	if err != nil {
		t.Error("can't be set limit")
	}

	_, err = operations.Upsert("", []byte("value"), 1)
	if err == nil {
		t.Error("key cannot be empty")
	}
}
