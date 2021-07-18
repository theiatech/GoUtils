package FileUtil

import "os"

func FileIsExists(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if nil == err {
		if fileInfo.IsDir() {
			return false
		}
		return true
	}
	return false
}
