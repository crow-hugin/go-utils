package format

import "strings"

func ParseQueryParamsToMap(s string) map[string]interface{} {
	arr := strings.Split(s, "&")
	ret := make(map[string]interface{})
	if len(s) < 2 {
		return ret
	}
	for _, value := range arr {
		vArr := strings.Split(value, "=")
		ret[vArr[0]] = vArr[1]
	}
	return ret
}
