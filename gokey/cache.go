package gokey

import (
	"errors"
	"time"
)

type Cache struct {
	pairsSet map[string]pair //contains expiration time and value of a key
}

type pair struct {
	ttl       time.Duration
	createdAt time.Time
	value     []byte
}

var (
	_ Operations = (*Cache)(nil)

	ErrEmptyKey       = errors.New("key cannot be empty")
	ErrNoExistKey     = errors.New("key does not exist")
	ErrNotImplemented = errors.New("not implemented")
)

// Get the values of the key, if this exists in the cache
func (c *Cache) Get(key string) ([]byte, error) {
	if isEmpty(key) {
		return nil, ErrEmptyKey
	}

	keyEncrypted := generateMD5HashFromKey([]byte(key))
	pair, exists := c.pairsSet[keyEncrypted]

	if !exists {
		return nil, ErrNoExistKey
	}

	if time.Since(pair.createdAt) > pair.ttl && pair.ttl != -1 {
		delete(c.pairsSet, keyEncrypted)
		return nil, ErrNoExistKey
	}

	return pair.value, nil
}

// Upsert cache a new key pair or update an existing one
// if ttl is equals to zero the key will not expire
func (c *Cache) Upsert(key string, value []byte, ttl time.Duration) (bool, error) {

	if isEmpty(key) {
		return false, ErrEmptyKey
	}

	var keyEncrypted = generateMD5HashFromKey([]byte(key))

	if c.pairsSet == nil {
		c.pairsSet = make(map[string]pair)
	}

	// redis in generic command:  if (ttl == -1)
	// golang use with functions time.Duration = -1
	c.pairsSet[keyEncrypted] = pair{
		ttl:       ttl,
		createdAt: time.Now(),
		value:     value,
	}

	return true, nil
}

func (c *Cache) Delete(key string) (bool, error) {
	if isEmpty(key) {
		return false, ErrEmptyKey
	}

	var keyEncrypted = generateMD5HashFromKey([]byte(key))

	_, exists := c.pairsSet[keyEncrypted]

	if exists {
		delete(c.pairsSet, keyEncrypted)
	} else {
		return false, errors.New("key not found")
	}

	return true, nil

}
