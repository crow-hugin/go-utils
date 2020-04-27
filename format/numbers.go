package format

import (
	"math/rand"
	"strconv"
	"time"
)

func StringToInt(str string) int {
	int, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return int
}

func StringToInt32(str string) int32 {
	num, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int32(num)
}

func StringToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func StringToFloat32(str string) float32 {
	num, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}
	return float32(num)
}

func StringToFloat64(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return num
}

// 生成随机数
func GenerateRandom(interval int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(interval)
}
