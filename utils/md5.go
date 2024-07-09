package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 md5加密
func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
