package gokey

import (
	"encoding/binary"
	"errors"
	"flag"
)

const defaultTupleMaxSize int = 1000 //bytes = 1 kb
const defaultPairsLimit int = 10000  //items mapping

var (
	tupleMaxSizeSetting = flag.Int("limit", defaultTupleMaxSize, "Limit tuple value size")
	limitPairsSetting   = flag.Int("pairs", defaultPairsLimit, "Limit map pairs set size")
)

func InitFlags() {
	flag.Parse()
}

// Set a custom limit by env arg or flag
// set default if not defined on start up
// default to 32
func getTupleMaxSize() int {
	// go run example.go -limit 50
	if *tupleMaxSizeSetting > 0 {
		return *tupleMaxSizeSetting
	}

	return defaultTupleMaxSize
}

func getLimitPairsSet() int {
	// go run example.go -pairs 1000
	if *limitPairsSetting > 0 {
		return *limitPairsSetting
	}

	return defaultPairsLimit
}

// Check cache value size to verify
// if not greater than the limit setting
func (c *Cache) checkTupleMaxSize(value []byte) error {
	// valTpl := []byte("32-byte-long-value")
	if binary.Size(value) > getTupleMaxSize() {
		return errors.New("cache tuple value size out of limit setting")
	}

	return nil
}

func (c *Cache) checkPairsSetLimit(pairsSet *map[string]tuple) error {
	if len(*pairsSet) > getLimitPairsSet() {
		return errors.New("cache map size out of limit setting")
	}

	return nil
}
