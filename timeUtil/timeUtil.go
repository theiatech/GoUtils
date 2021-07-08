package timeUtil

import (
	"time"
)

func Year() string {
	return time.Now().Format("2006")
}

func Month() string {
	return time.Now().Format("01")
}

func Day() string {
	return time.Now().Format("02")
}

func Hour() string {
	return time.Now().Format("15")
}

func Minute() string {
	return time.Now().Format("04")
}

func Second() string {
	return time.Now().Format("05")
}

// 例如 2006-01-02
func Date() string {
	return time.Now().Format("2006-01-02")
}

// 例如 15:04:05
func Time() string {
	return time.Now().Format("15:04:05")
}

// 例如 2006-01-02 15:04:05
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 例如 1616411595
func Timestamp() int64 {
	return time.Now().Unix()
}

// 例如 1616411595 -> 2021-03-22 19:13:15
func Timestamp2String(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

/**
* timeLayout 转化所需模板 2006-01-02 15:04:05
* toBeCharge 待转化为时间戳的字符串 2015-01-01 00:00:00
 */
func String2Timestamp(timeLayout, toBeCharge string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc)
	return theTime.Unix()
}
