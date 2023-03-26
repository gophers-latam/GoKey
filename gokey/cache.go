package gokey

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	pairsSet map[string]tuple //contains expiration time and value of a key

	hashFn func([]byte) string
}

type tuple struct {
	ttl       time.Duration
	createdAt time.Time
	value     []byte
}

type Options struct {
	MaxSize int
	AHast   string
	TTL     float64 // in Newcache or their methods?
}

var (
	_ Operations = (*Cache)(nil)

	ErrEmptyKey   = errors.New("key cannot be empty")
	ErrNoExistKey = errors.New("key does not exist")
	ErrExpiredKey = errors.New("key has expired")
)

func newCache(o ...*Options) *Cache {
	var options *Options

	if len(o) < 1 || o[0] == nil {
		options = &Options{}
	} else {
		options = o[0]
	}

	hashFn := selectHash(options.AHast)
	return &Cache{
		RWMutex:  sync.RWMutex{},
		pairsSet: make(map[string]tuple, sizeLimit(options.MaxSize)),
		hashFn:   hashFn,
	}
}

// Get the values of the key, if this exists in the cache
func (c *Cache) Get(key string) ([]byte, error) {
	if isEmpty(key) {
		return nil, ErrEmptyKey
	}

	c.RLock()
	defer c.RUnlock()

	keyHashed := c.hashFn([]byte(key))

	pair, exists := c.pairsSet[keyHashed]

	if !exists {
		return nil, ErrNoExistKey
	}

	if time.Since(pair.createdAt) > pair.ttl && pair.ttl != -1 {
		delete(c.pairsSet, keyHashed)
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

	errPairs := c.checkPairsSetLimit(&c.pairsSet)
	if errPairs != nil {
		return false, errPairs
	}

	errTuple := c.checkTupleMaxSize(value)
	if errTuple != nil {
		return false, errTuple
	}

	c.Lock()
	defer c.Unlock()

	keyHashed := c.hashFn([]byte(key))

	if c.pairsSet == nil {
		c.pairsSet = make(map[string]tuple, getLimitPairsSet())
	}

	// redis in generic command:  if (ttl == -1)
	// golang use with functions time.Duration = -1
	c.pairsSet[keyHashed] = tuple{
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

	c.Lock()
	defer c.Unlock()

	keyHashed := c.hashFn([]byte(key))

	_, exists := c.pairsSet[keyHashed]

	if exists {
		delete(c.pairsSet, keyHashed)
	} else {
		return false, errors.New("key not found")
	}

	return true, nil
}

func (c *Cache) Exists(key string) (bool, error) {
	if isEmpty(key) {
		return false, ErrEmptyKey
	}

	c.RLock()
	defer c.RUnlock()

	keyHashed := c.hashFn([]byte(key))

	pair, exists := c.pairsSet[keyHashed]

	if !exists {
		return false, ErrNoExistKey
	}

	if time.Since(pair.createdAt) > pair.ttl && pair.ttl != -1 {
		return false, ErrExpiredKey
	}

	return true, nil
}
