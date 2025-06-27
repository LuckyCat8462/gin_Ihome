package model

import (
	"crypto/md5"
	"encoding/hex"
)

func RegisterUser(mobile, pwd string) error {
	var user User
	user.Name = mobile // 默认使用手机号作为用户名

	// 使用 md5 对 pwd 加密
	m5 := md5.New()                             // 初始md5对象
	m5.Write([]byte(pwd))                       // 将 pwd 写入缓冲区
	pwd_hash := hex.EncodeToString(m5.Sum(nil)) // 不使用额外的秘钥

	user.Password_hash = pwd_hash

	// 插入数据到MySQL
	return GlobalConn.Create(&user).Error
}
