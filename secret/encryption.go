package secret

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"strings"
)

// Sha256Encryption 密码加密
func Sha256Encryption(str string) string {
	sh := sha256.New()
	sh.Write([]byte(str))
	return hex.EncodeToString(sh.Sum([]byte("app")))
}

// MakeUUID 生成uuid
func MakeUUID() string {
	return uuid.Must(uuid.NewUUID()).String()
}

// MakeEncryptionUUID 生成加密后的UUID
func MakeEncryptionUUID(str string) string {
	return uuid.NewSHA1(uuid.New(), []byte(str)).String()
}

// MakeDbID 生成数据ID
func MakeDbID() string {
	return strings.Replace(MakeUUID(), "-", "", -1)
}

// Md5Encode 小写的
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// MD5Encode 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// MakePassword 创建密码
func MakePassword(enterPassword, salt string) string {
	return Md5Encode(enterPassword + salt)
}

// ValidatePassword 验证密码
func ValidatePassword(enterPassword, salt, password string) bool {
	return Md5Encode(enterPassword+salt) == password
}
