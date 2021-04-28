package gokey

import "time"

// Operations contains all the basic operations for all the interactions with the data structure (cache).
type Operations interface {
	// get returns a value and an optional error given a key.
	get(key string) (interface{}, error)

	// upsert is for create/update operation in the cache. If 0 or negative, the entry will not expire.
	// Returns whether the entry was created with this operation or not (updated) and an optional error
	upsert(key string, value interface{}, ttl time.Duration) (bool, error)

	// delete removes a value given a key.
	// Returns whether the entry was deleted or not and an optional error
	delete(key string) (bool, error)
}
