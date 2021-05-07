package tests

import (
	"testing"

	"github.com/gophers-latam/GoKey/gokey"
)

// go test -run TestCacheUpsert -v
func TestCacheUpsert(t *testing.T) {
	var operations gokey.Operations = new(gokey.Cache)

	_, err := operations.Upsert("key", []byte("value"), -1)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = operations.Upsert("", []byte("value"), 1)
	if err != nil {
		t.Error(err.Error())
	}

	// case ok
	_, err = operations.Upsert("key", []byte("value"), 0)
	if err != nil {
		t.Error("here: unexpected error")
	}
}
