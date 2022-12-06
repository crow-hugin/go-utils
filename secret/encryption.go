package secret

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/google/uuid"
	"hash"
	"strings"
)

type HmacCryptoFun func(string, string) string

type CryptoFun func(string) string

// MakeUUID 生成uuid
func MakeUUID() string {
	return uuid.Must(uuid.NewUUID()).String()
}

// MakeEncryptionUUID 生成加密后的UUID
func MakeEncryptionUUID(str string) string {
	return uuid.NewSHA1(uuid.New(), []byte(str)).String()
}

// MakeTrimUUID 生成去除-后的uuid
func MakeTrimUUID() string {
	return strings.Replace(MakeUUID(), "-", "", -1)
}

// MD5Upper 大写MD5
func MD5Upper(data string) string {
	return strings.ToUpper(MD5(data))
}

// MD5Double 两次MD5
func MD5Double(data string) string {
	return MD5(MD5(data))
}

func HmacSHA512(data, key string) string {
	return common(hmac.New(sha512.New, []byte(key)), data)
}

func SHA512(data string) string {
	return common(sha512.New(), data)
}

func HmacSHA384(data, key string) string {
	return common(hmac.New(sha512.New384, []byte(key)), data)
}

func SHA384(data string) string {
	return common(sha512.New384(), data)
}

func HmacSHA256(data, key string) string {
	return common(hmac.New(sha256.New, []byte(key)), data)
}

func SHA256(data string) string {
	return common(sha256.New(), data)
}

func HmacSHA224(data, key string) string {
	return common(hmac.New(sha256.New224, []byte(key)), data)
}

func SHA224(data string) string {
	return common(sha256.New224(), data)
}

func HmacMD5(data, key string) string {
	return common(hmac.New(md5.New, []byte(key)), data)
}

func MD5(data string) string {
	return common(md5.New(), data)
}

func common(h hash.Hash, data string) string {
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
