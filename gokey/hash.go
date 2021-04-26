package gokey

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateHashFromKey(key string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(key))
	return hex.EncodeToString(algorithm.Sum(nil))
}
