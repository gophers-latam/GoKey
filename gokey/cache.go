package gokey

import (
	"errors"
	"time"
)

type Cache struct {
	pairsSet map[string]pair //contains expiration time and value of a key
}

type pair struct {
	ttl   time.Duration
	value []byte
}

var (
	_             Operations = (*Cache)(nil)
	ErrorEmptyKey            = errors.New("key cannot be empty")
)

// Get the values of the key, if this exists in the cache
func (c *Cache) Get(key string) ([]byte, error) {
	if isEmpty(key) {
		return nil, ErrorEmptyKey
	}

	keyEncrypted := generateMD5HashFromKey([]byte(key))
	pair, exists := c.pairsSet[keyEncrypted]

	if exists {
		return pair.value, nil
	}

	return nil, errors.New("key has no related values")
}

// Upsert cache a new key pair or update an existing one
// if ttl is equals to zero the key will not expire.
func (c *Cache) Upsert(key string, value []byte, ttl time.Duration) (bool, error) {

	if isEmpty(key) {
		return false, ErrorEmptyKey
	}

	var keyEncrypted = generateMD5HashFromKey([]byte(key))

	// redis in generic command:  if (ttl == -1)
	// golang use with functions time.Duration = -1
	c.pairsSet[keyEncrypted] = pair{
		ttl:   ttl,
		value: value,
	}

	return true, nil
}

// Delete the key in the shared structure.
func (c *Cache) Delete(key string) (bool, error) {
	return false, errors.New("not implemented")
}
