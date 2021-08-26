package FileUtil

import (
	"fmt"
	"os"
)

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

func AppendToFile(filename string,data []byte) bool {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", filename)
		return false
	}
	defer f.Close()
	_,err = f.Write(data)
	return nil == err
}

func AppendStringToFile(filename ,str string) bool {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", filename)
		return false
	}
	defer f.Close()
	_,err = f.WriteString(str)
	return nil == err
}