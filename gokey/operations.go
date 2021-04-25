package gokey

// Operations contains all the basic operations for all the interactions with the data structure (cache).
type Operations interface {
	// get a value given a key.
	get(key string) (string, error)
	// upsert is for create/update operation in the cache. ttl should be expressed in millis.
	// If 0, the entry will not expire.
	upsert(key, value string, ttl int) (bool, error)
	//delete a value given a key
	delete(key string) (bool, error)
}
