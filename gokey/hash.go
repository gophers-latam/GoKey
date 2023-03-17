package gokey

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
)

// generateMD5 is a hash generator function according to input(key)
// using md5 algorithm.
func generateMD5(key []byte) (string, error) {
	algorithm := md5.New()
	_, err := algorithm.Write(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(algorithm.Sum(nil)), nil
}

// generateSHA512
func generateSHA512(input string) (string, error) {
	hasher := sha512.New()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		return "", err
	}
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash), nil
}
