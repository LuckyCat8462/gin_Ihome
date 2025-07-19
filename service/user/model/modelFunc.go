package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 创建全局变量redis,连接池句柄
var RedisPool redis.Pool

// 创建一个函数,用于初始化连接池
func InitRedis() {
	//链接redis
	RedisPool = redis.Pool{
		MaxIdle:         20,
		MaxActive:       50,
		MaxConnLifetime: 60 * 5,
		IdleTimeout:     60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.81.128:6379")
		},
	}
	//RedisPool = redis.Pool{
	//	MaxIdle:         20,
	//	MaxActive:       50,
	//	MaxConnLifetime: 60 * 5,
	//	IdleTimeout:     60,
	//	Dial: func() (redis.Conn, error) {
	//		return redis.Dial("tcp", "192.168.81.128:6379")
	//	},
	//}
}
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

// // 校验图片验证码
func CheckImgCode(uuid, imgCode string) bool {
	// 链接 redis --- 从链接池中获取链接
	conn, err := redis.Dial("tcp", "192.168.81.128:6379")
	if err != nil {
		fmt.Println("redis.Dial err:", err)
		return false
	}
	//conn := RedisPool.Get()
	defer conn.Close()

	// 查询 redis 数据
	code, err := redis.String(conn.Do("get", uuid))
	if err != nil {
		fmt.Println("查询错误 err:", err)
		return false
	}

	// 返回校验结果
	return code == imgCode
}

func SaveRealName(userName, realName, idCard string) error {
	err := GlobalConn.Model(&User{}).Where("name = ?", userName).
		Update("id_card", idCard)
	fmt.Println("save情况", err)
	return nil

	//Updates(map[string]interface{}{"real_name": realName, "id_card": idCard}).Error

}
