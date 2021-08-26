package CryptoUtil

import (
	"crypto/md5"
	"fmt"
)

func Md5Encrypt(s string) string {
	return fmt.Sprintf("%x",md5.Sum([]byte(s)))
}
