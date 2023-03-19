package gokey

import (
	"encoding/binary"
	"errors"
	"flag"
)

const defaultTupleLimit int = 32   //bytes
const defaultPairsLimit int = 1000 //items

var (
	limitTupleSetting = flag.Int("limit", defaultTupleLimit, "Limit tuple value size")
	limitPairsSetting = flag.Int("pairs", defaultPairsLimit, "Limit map pairs set size")
)

// Set a custom limit by env arg or flag
// set default if not defined on start up
// default to 32
func getLimitTupleValue() int {
	// go run example.go -limit 50
	flag.Parse()
	if *limitTupleSetting > 0 {
		return *limitTupleSetting
	} else {
		return defaultTupleLimit
	}
}

func getLimitPairsSet() int {
	// go run example.go -pairs 1000
	flag.Parse()
	if *limitPairsSetting > 0 {
		return *limitPairsSetting
	} else {
		return defaultPairsLimit
	}
}

// Check cache value size to verify
// if not greater than the limit setting
func checkTupleLimit(value []byte) error {
	// valTpl := []byte("32-byte-long-value")
	if binary.Size(value) > getLimitTupleValue() {
		return errors.New("cache tuple value size out of limit setting")
	}

	return nil
}

func checkPairsSetLimit(pairsSet *map[string]tuple) error {
	if len(*pairsSet) > getLimitPairsSet() {
		return errors.New("cache map size out of limit setting")
	}

	return nil
}
