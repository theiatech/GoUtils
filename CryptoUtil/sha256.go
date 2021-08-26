package CryptoUtil

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Encrypt(s string) string {
	return fmt.Sprintf("%x",sha256.Sum256([]byte(s)))
}
