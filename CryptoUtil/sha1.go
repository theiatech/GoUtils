package CryptoUtil

import (
	"crypto/sha1"
)

func Sha1(b []byte) []byte {
	r := sha1.Sum(b)
	return r[:]
}
