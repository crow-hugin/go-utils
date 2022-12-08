package format

import (
	"strings"
	"time"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
)

// GetNowTime 获取当前日期
func GetNowTime(format ...string) string {
	if len(format) == 0 {
		format = append(format, DateTimeFormat)
	}
	t := time.Unix(time.Now().Unix(), 0).Format(format[0])
	return t
}

// GetNowUnix 获取当前时间
func GetNowUnix() int64 {
	return time.Now().Unix()
}

// GetUnixToday 获取当天几号
func GetUnixToday(timestamp int64) int {
	return time.Unix(timestamp, 0).Day()
}

// GetNowDate 获取当前日期
func GetNowDate() string {
	t := time.Unix(GetNowUnix(), 0).Format(DateFormat)
	return t
}

func GetUnixDateTimestamp(timestamp int64) int64 {
	var dateTime = time.Unix(timestamp, 0)
	tm := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, dateTime.Location())
	return tm.UnixMilli()
}

// GetNowHourUnix 获取当前小时时间
func GetNowHourUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	return tm.Unix()
}

// GetZeroHourUnix 获取零时时间
func GetZeroHourUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

// GetEarlyMonthUnix 获取月初时间
func GetEarlyMonthUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

// GetEarlyYearUnix 获取年初时间
func GetEarlyYearUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

// GetUnixToFormatString 将时间戳转换成字符串
func GetUnixToFormatString(timestamp int64, f string) string {
	return time.Unix(timestamp, 0).Format(f)
}

// GetUnixToString 时间戳转换字符串
func GetUnixToString(timestamp int64, format ...string) string {
	if len(format) == 0 {
		format = append(format, DateTimeFormat)
	}
	return GetUnixToFormatString(timestamp, format[0])
}

// GetUnixToHourString 小时时间戳转换字符串
func GetUnixToHourString(timestamp int64) string {
	return GetUnixToFormatString(timestamp, "15:04")
}

// GetUnixToOldYearTime 往前多少年时间戳
func GetUnixToOldYearTime(i int) int64 {
	currentYear := time.Now().Year()
	oldMonth := currentYear - i
	t := time.Date(oldMonth, 1, 1, 0, 0, 0, 0, time.Local)
	return t.Unix()
}

func GetBeginToEndDate(begin, end time.Time) (days []time.Time, err error) {
	dd := end.Sub(begin)
	for i := 0; i <= int(dd.Hours()/24); i++ {
		d, _ := time.ParseDuration("-24h")
		day := end.Add(d * time.Duration(i))
		days = append(days, day)
	}
	return
}

func GetBeginToEndDateStr(begin, end string) (days []string, err error) {
	var beginTime, endTime time.Time

	beginTime, err = time.Parse(DateFormat, begin)
	if err != nil {
		return nil, err
	}

	endTime, err = time.Parse(DateFormat, end)
	if err != nil {
		return nil, err
	}
	dayTimes, err := GetBeginToEndDate(beginTime, endTime)
	if err != nil {
		return nil, err
	}
	for _, v := range dayTimes {
		days = append(days, time.Unix(v.Unix(), 0).Format(DateFormat))
	}
	return
}

func GetDateToTimestamp(dateTime string) int64 {
	t, err := time.Parse(DateFormat, dateTime)
	if err != nil {
		return 0
	}
	return t.UnixMilli()
}

func GetDateTimeToTimestamp(dateTime string) int64 {
	t, err := time.Parse(DateTimeFormat, dateTime)
	if err != nil {
		return 0
	}
	return t.UnixMilli()
}

// AddTimeHours 时间加减
func AddTimeHours(date, hour string) string {
	atime, _ := time.Parse(DateTimeFormat, date)
	d, _ := time.ParseDuration(hour)
	return atime.Add(d).Format(DateTimeFormat)
}

// HisToSecond 时间转秒
func HisToSecond(t string) int {
	st := strings.Split(t, ":")
	h := StringToInt(st[0])
	m := StringToInt(st[1])
	s := StringToInt(st[2])
	return 3600*h + 60*m + s
}

func HisAddSec(t, sec string) string {
	s, _ := time.Parse(TimeFormat, t)
	m, _ := time.ParseDuration(sec)
	return s.Add(m).Format(TimeFormat)
}
