package utils

import "time"

// 获取当前时间
func GetNowUnix() int64 {
	return time.Now().Unix()
}

// 获取当天几号
func GetUnixToday(timestamp int64) int {
	return time.Unix(timestamp, 0).Day()
}

// 获取当前小时时间
func GetNowHourUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	return tm.Unix()
}

// 获取零时时间
func GetZeroHourUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

// 获取月初时间
func GetEarlyMonthUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

// 获取年初时间
func GetEarlyYearUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

// 将时间戳转换成字符串
func GetUnixToFormatString(timestamp int64, f string) string {
	return time.Unix(timestamp, 0).Format(f)
}

// 时间戳转换字符串
func GetUnixToString(timestamp int64) string {
	return GetUnixToFormatString(timestamp, "2006-01-02 15:04:05")
}

// 小时时间戳转换字符串
func GetUnixToHourString(timestamp int64) string {
	return GetUnixToFormatString(timestamp, "15:04")
}

//往前多少年时间戳
func GetUnixToOldYearTime(i int) int64 {
	currentYear := time.Now().Year()
	oldMonth := currentYear - i
	t := time.Date(oldMonth, 1, 1, 0, 0, 0, 0, time.Local)
	return t.Unix()
}

//往前多少天时间戳
func GetUnixToOldTimeDay(i int) int64 {
	day := time.Now().Day()
	oldMonth := day - i
	t := time.Date(time.Now().Year(), time.Now().Month(), oldMonth, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	return t.Unix()
}
