package FileUtil

import "os"

func isExists(filename string) bool {
	_, err := os.Stat(filename)
	if nil == err {
		return true
	}
	return false
}
