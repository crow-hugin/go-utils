package format

import (
	"math/rand"
	"strconv"
	"time"
)

// StringToInt 字符串转int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// StringToInt32 字符串转int32
func StringToInt32(str string) int32 {
	num, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int32(num)
}

// StringToInt64 字符串转int64
func StringToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

// StringToFloat32 字符串转int32
func StringToFloat32(str string) float32 {
	num, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}
	return float32(num)
}

// StringToFloat64 字符串转float64
func StringToFloat64(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return num
}

// GenRandom 生成随机数
func GenRandom(interval int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(interval)
}
