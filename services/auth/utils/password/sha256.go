package password

import (
	"crypto/sha256"
	"encoding/hex"
)

const Salt = "salt-lijialnag666-2024-11-12" // TODO : 目前使用静态加盐

func Encrypt(password string) string {
	// 密码加盐
	hash := sha256.New()
	hash.Write([]byte(password + Salt)) // TODO: 写入配置文件
	return hex.EncodeToString(hash.Sum(nil))
}

// 动态加盐存入数据库
//func GenerateSalt(length int) (string, error) {
//	salt := make([]byte, length)
//	if _, err := rand.Read(salt); err != nil {
//		return "", err
//	}
//	return hex.EncodeToString(salt), nil
//}
//
//func encrypt(password, salt string) string {
//	hash := sha256.New()
//	hash.Write([]byte(password + salt))
//	return hex.EncodeToString(hash.Sum(nil))
//}
