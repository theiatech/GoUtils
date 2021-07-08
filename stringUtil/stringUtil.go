package stringUtil

import (
	"math/rand"
	"regexp"
	"time"
)

func RandomString(l int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func FindStr(src, expr string) string {
	compile, err := regexp.Compile(expr)
	if nil != err {
		return ""
	}
	result := compile.FindAllStringSubmatch(src, 1)
	for _, data := range result {
		return data[0]
	}
	return ""
}

func FindAllStr(src, expr string) []string {
	r := []string{}
	compile, err := regexp.Compile(expr)
	if nil != err {
		return r
	}
	results := compile.FindAllStringSubmatch(src, 1)
	for _, result := range results {
		r = append(r, result[1])
	}
	return r
}
