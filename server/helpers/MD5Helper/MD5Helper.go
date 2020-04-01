package MD5Helper

import (
	"crypto/md5"
	"fmt"
)

// MD5 encryption（32bit）
func EncryptMD5To32Bit(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
