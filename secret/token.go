package secret

import (
	"encoding/base64"
	"strings"
)

const Key = "crow_hugin@16888"

func EncryptToken(s string) string {
	result, err := AesEncrypt([]byte(s), []byte(Key))
	if err != nil {
		return ""
	}
	ss := base64.StdEncoding.EncodeToString(result)
	ss = strings.Replace(ss, "/", "_a", -1)
	ss = strings.Replace(ss, "+", "_b", -1)
	ss = strings.Replace(ss, "=", "_c", -1)
	return ss
}

func DecryptToken(s string) string {
	s = strings.Replace(s, "_a", "/", -1)
	s = strings.Replace(s, "_b", "+", -1)
	s = strings.Replace(s, "_c", "=", -1)

	ss, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	result, err := AesDecrypt(ss, []byte(Key))
	if err != nil {
		return ""
	}
	return string(result)
}
