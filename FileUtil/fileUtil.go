package FileUtil

import "os"

func IsExists(filename string) bool {
	_, err := os.Stat(filename)
	if nil == err {
		return true
	}
	return false
}
