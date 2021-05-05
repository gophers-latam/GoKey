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

//implementation cheack on compilation time
var (
	_ Operations = (*Cache)(nil)

	ErrorEmptyKey = errors.New("The key cannot be empty")
)

func (this *Cache) Get(key string) ([]byte, error) {
	if isEmpty(key) {
		return nil, ErrorEmptyKey
	}

	return []byte(""), errors.New("not implemented")
}

func (this *Cache) Upsert(key string, value []byte, ttl time.Duration) (bool, error) {

	if isEmpty(key) {
		return false, ErrorEmptyKey
	}

	var keyEncrypted string = generateMD5HashFromKey([]byte(key))

	if this.pairsSet == nil {
		this.pairsSet = make(map[string]pair)
	}

	if ttl < 0 {
		return false, errors.New("Ttl cannot be lower than 0")

	} else if ttl > 0 {
		time.AfterFunc(time.Duration(ttl)*time.Millisecond, func() {
			delete(this.pairsSet, keyEncrypted)
		})

	} else {
		//if ttl is equals to zero the key will not expire
		ttl = -1
	}

	this.pairsSet[keyEncrypted] = pair{
		ttl:   ttl,
		value: []byte(value),
	}

	return true, nil
}

func (this *Cache) Delete(key string) (bool, error) {
	return false, errors.New("not implemented")
}

func isEmpty(key string) bool {
	return key == ""
}
