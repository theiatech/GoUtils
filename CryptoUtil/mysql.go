package CryptoUtil

import "fmt"

func MysqlPassword(b []byte) string {
	return fmt.Sprintf("*%x",Sha1(Sha1(b)))
}

