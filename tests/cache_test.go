package tests

import (
	"testing"

	"github.com/gophers-latam/GoKey/gokey"
)

func TestCacheUpsert(t *testing.T) {

	var operations gokey.Operations = new(gokey.Cache)
	_, err := operations.Upsert("key", []byte("value"), -1)
	if err == nil {
		t.Error("the ttl argument doesn't accept negative numbers.")
	}
	_, err = operations.Upsert("", []byte("value"), -1)

	if err == nil {
		t.Error("the key argument doesn't accept empty string.")
	}

}
