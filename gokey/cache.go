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
	_ Operations = (*Cache)(nil)

	ErrorEmptyKey = errors.New("key cannot be empty")
)

// Get the values of the key, if this exists in the cache
func (this *Cache) Get(key string) ([]byte, error) {
	if isEmpty(key) {
		return nil, ErrorEmptyKey
	}

	keyEncrypted := generateMD5HashFromKey([]byte(key))
	pair, exists := this.pairsSet[keyEncrypted]

	if exists {
		return pair.value, nil
	}

	return nil, errors.New("key has no related values")
}

// Upsert cache a new key pair or update an existing one
// if ttl is equals to zero the key will not expire
func (this *Cache) Upsert(key string, value []byte, ttl time.Duration) (bool, error) {

	if isEmpty(key) {
		return false, ErrorEmptyKey
	}

	var keyEncrypted string = generateMD5HashFromKey([]byte(key))

	if this.pairsSet == nil {
		this.pairsSet = make(map[string]pair)
	}

	if ttl < 0 {
		// redis is: if (ttl < 0) ttl = 0;
		return false, errors.New("ttl value cannot be lower than 0")

	} else if ttl > 0 {
		time.AfterFunc(time.Duration(ttl)*time.Millisecond, func() {
			delete(this.pairsSet, keyEncrypted)
		})

	} else {
		ttl = -1
	}
	// redis in generic command:  if (ttl == -1)
	// golang use with functions time.Duration = -1

	this.pairsSet[keyEncrypted] = pair{
		ttl:   ttl,
		value: []byte(value),
	}

	return true, nil
}

func (this *Cache) Delete(key string) (bool, error) {
	return false, errors.New("not implemented")
}
