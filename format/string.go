package format

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const (
	KcRandKindNum   = 0 // 纯数字
	KcRandKindLower = 1 // 小写字母
	KcRandKindUpper = 2 // 大写字母
	KcRandKindAll   = 3 // 数字、大小写字母
)

func IsNil(s string) bool {
	if s == "" || len(s) == 0 {
		return true
	}
	return false
}
func IsNotNil(s string) bool {
	return !IsNil(s)
}

func IsInt(s string) bool {
	if IsNil(s) {
		return false
	}
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

// Sub 截取字符串 start 起点下标 end 终点下标(不包括)
func Sub(str string, start, end int) (string, error) {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return "", errors.New("start is wrong")
	}

	if end < 0 || end > length {
		return "", errors.New("end is wrong")
	}
	return string(rs[start:end]), nil
}

// RandString 随机字符串
// KcRandKindNum   = 0 // 纯数字
// KcRandKindLower = 1 // 小写字母
// KcRandKindUpper = 2 // 大写字母
// KcRandKindAll   = 3 // 数字、大小写字母
func RandString(size int, kind int) []byte {
	k, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if all { // random k
			k = rand.Intn(3)
		}
		scope, base := kinds[k][0], kinds[k][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

func ToString(target interface{}) string {
	switch t := target.(type) {
	case string:
		return t
	case float32, float64:
		return fmt.Sprintf("%f", t)
	case int, int16, int32, int64:
		return fmt.Sprintf("%d", t)
	case time.Time:
		return fmt.Sprintf("%s", t)
	default:
		fmt.Printf("%s", t)
	}
	return ""
}

func StrSliceToInt(data []string, t int) ([]int, []string) {
	var ins []int
	var str []string

	if len(data) == 0 {
		return nil, nil
	}

	for _, r := range data {
		ins = append(ins, StringToInt(r))
	}
	// 返回没有排序的数据
	if t == 1 {
		return ins, nil
	}

	sort.Ints(ins)
	for _, ri := range ins {
		str = append(str, strconv.Itoa(ri))
	}
	return ins, str
}
