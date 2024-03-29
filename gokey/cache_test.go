package gokey

import (
	"errors"
	"strconv"
	"sync"
	"testing"
	"time"
)

var operations = newCache(&Options{AHast: MD5})

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

// go test -run TestCacheConcurrentUpsert -v
func TestCacheConcurrentUpsert(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key" + strconv.Itoa(i)
			value := "value" + strconv.Itoa(i)
			operations.Upsert(key, []byte(value), -1)
		}(i)
	}

	wg.Wait()

	value, err := operations.Get("key1")
	if err != nil {
		t.Error(err.Error())
	}

	value2, err2 := operations.Get("key2")
	if err2 != nil {
		t.Error(err.Error())
	}

	res := string(value)
	res2 := string(value2)

	if res != "value1" || res2 != "value2" {
		t.Error("error while concurrently accessing in the cache")
	}
}

// go test -run TestCacheGetExpiredKey -v
func TestCacheGetExpiredKey(t *testing.T) {
	_, err := operations.Upsert("key", []byte("value"), 100*time.Millisecond)
	if err != nil {
		t.Error("expected no errors in Upsert method, got:", err.Error())
	}
	time.Sleep(100 * time.Millisecond)

	_, err = operations.Get("key")
	if err == nil {
		t.Error("expected ErrExpiredKey, got: nil")
	}
	if err != nil {
		if !errors.Is(err, ErrNoExistKey) {
			t.Error("expected ErrExpiredKey, got:", err.Error())
		}
	}

	_, err = operations.Get("key")
	if err != nil {
		if !errors.Is(err, ErrNoExistKey) {
			t.Error("expected 'key does not exist' error message, got:", err.Error())
		}
	}
}

// go test -run TestCacheConcurrentGet -v
func TestCacheConcurrentGet(t *testing.T) {
	var wg sync.WaitGroup

	_, err := operations.Upsert("key1", []byte("value1"), 10*time.Second)
	if err != nil {
		t.Error("expected no errors in Upsert method, got:", err.Error())
	}

	_, err = operations.Upsert("key2", []byte("value2"), 10*time.Second)
	if err != nil {
		t.Error("expected no errors in Upsert method, got:", err.Error())
	}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key" + strconv.Itoa(i)
			value, err := operations.Get(key)

			t.Log(string(value))

			if err != nil {
				t.Error("expected no errors in Get method, got:", err.Error())
			}

			if value == nil {
				t.Error("expected a value, got nil")
			}
		}(i)
	}

	wg.Wait()
}

// go test -run TestCacheDelete -v
func TestCacheDelete(t *testing.T) {
	_, err := operations.Upsert("key", []byte("value"), 10)
	if err != nil {
		t.Error("Expected no errors in Upsert method, got:", err.Error())
	}

	_, err = operations.Delete("key")
	if err != nil {
		t.Error("Expected no errors in Delete method, got:", err.Error())
	}
}

// go test -run TestCacheConcurrentDelete -v
func TestCacheConcurrentDelete(t *testing.T) {
	var wg sync.WaitGroup

	_, err := operations.Upsert("key1", []byte("value1"), 10*time.Second)
	if err != nil {
		t.Error("expected no errors in Upsert method, got:", err.Error())
	}

	_, err = operations.Upsert("key2", []byte("value2"), 10*time.Second)
	if err != nil {
		t.Error("expected no errors in Upsert method, got:", err.Error())
	}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key" + strconv.Itoa(i)
			_, err = operations.Delete(key)
			if err != nil {
				t.Error("Expected no errors in Delete method, got:", err.Error())
			}
		}(i)
	}

	wg.Wait()

	for i := 1; i <= 2; i++ {
		value, err := operations.Get("key" + strconv.Itoa(i))

		if value != nil {
			t.Error("expected nil, got:", err.Error())
		}
	}
}

// go test -run TestCacheExistsSomeKey -v
func TestCacheExistsSomeKey(t *testing.T) {

	_, err := operations.Upsert("key", []byte("value"), 10*time.Second)

	if err != nil {
		t.Errorf("expected no errors in Upsert method, got: %v", err.Error())
	}

	exists, err := operations.Exists("key")

	if err != nil {
		t.Errorf("expected no errors in Exists method, got: %v", err.Error())
	}

	if !exists {
		t.Errorf("expected it true, got: %v", exists)
	}

}

// go test -run TestCacheExistsExpiredKey -v
func TestCacheExistsExpiredKey(t *testing.T) {

	_, err := operations.Upsert("key", []byte("value"), 100*time.Millisecond)

	if err != nil {
		t.Errorf("expected no errors in Upsert method, got: %v", err.Error())
	}

	time.Sleep(100 * time.Millisecond)

	_, err1 := operations.Exists("key")

	if err1 != nil {
		if !errors.Is(err1, ErrExpiredKey) {
			t.Errorf("Ok. expected %v, got: %v", ErrExpiredKey, err1.Error())
		}
	}
}

// go test -run TestCacheExistsEmptyKey -v
func TestCacheExistsEmptyKey(t *testing.T) {
	_, err := operations.Exists("")
	if err != nil {
		if !errors.Is(err, ErrEmptyKey) {
			t.Errorf("expected '%v' error message, got %v", ErrEmptyKey, err.Error())
		}
	}
}

// go test -run TestCacheConcurrentExists -v
func TestCacheConcurrentExists(t *testing.T) {
	var wg sync.WaitGroup

	_, err := operations.Upsert("key1", []byte("value1"), 10*time.Second)
	if err != nil {
		t.Error("expected no errors in Upsert method, got:", err.Error())
	}

	_, err = operations.Upsert("key2", []byte("value2"), 10*time.Second)
	if err != nil {
		t.Error("expected no errors in Upsert method, got:", err.Error())
	}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key" + strconv.Itoa(i)
			exists, err := operations.Exists(key)

			if err != nil {
				t.Errorf("expected no errors in Exists method, got: %v", err.Error())
			}

			if !exists {
				t.Errorf("expected it true, got: %v", exists)
			}
		}(i)
	}

	wg.Wait()

	for i := 1; i <= 2; i++ {
		value, _ := operations.Get("key" + strconv.Itoa(i))

		if value == nil {
			t.Error("expected a value, got nil")
		}
	}
}
