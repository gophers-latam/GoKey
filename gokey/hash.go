package gokey

import (
	"crypto/md5"
	"encoding/hex"
)

//generateMD5HashFromKey is a hash generator function according to input(key)
//using md5 algorith
func generateMD5HashFromKey(key string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(key))
	return hex.EncodeToString(algorithm.Sum(nil))
}
