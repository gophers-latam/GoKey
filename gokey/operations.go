package gokey

import "time"

// Operations contains all the basic operations for all the interactions with the data structure (cache).
type Operations interface {
	// Get returns a value and an optional error given a key.
	Get(key string) ([]byte, error)

	// Upsert is for create/update operation in the cache. If ttl 0, the entry will not expire.
	// Returns whether the entry was created with this operation or not (updated) and an optional error
	Upsert(key string, value []byte, ttl time.Duration) (bool, error)

	// Delete removes a value given a key.
	// Returns whether the entry was deleted or not and an optional error
	Delete(key string) (bool, error)

	// Exists checks if a key is registered.
	// Returns true if it exists, false if not, and a optional error.
	Exists(key string) (bool, error)
}
