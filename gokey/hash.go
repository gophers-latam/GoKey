package gokey

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type closureHash func([]byte) (string, error)

// generateMD5 is a hash generator function according to input(key)
// using md5 algorithm.

func selectHash(tHash string) closureHash {
	switch tHash {
	case "sha256":
		return generateFromHash(sha256.New())
	case "sha1":
		return generateFromHash(sha1.New())
	default:
		return generateFromHash(md5.New())
	}
}

func generateFromHash(algorithm hash.Hash) closureHash {

	return func(key []byte) (string, error) {
		algorithm.Reset()
		// write never return error
		algorithm.Write(key)

		return hex.EncodeToString(algorithm.Sum(nil)), nil
	}
}
