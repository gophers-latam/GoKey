package gokey

import (
	"crypto/md5"
	"encoding/hex"
)

//generateMD5HashFromKey is a hash generator function according to input(key)
//using md5 algorithm
func generateMD5HashFromKey(key []byte) string {
	algorithm := md5.New()
	algorithm.Write(key)
	return hex.EncodeToString(algorithm.Sum(nil))
}
