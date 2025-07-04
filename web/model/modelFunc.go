package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 创建全局redis 连接池 句柄
var RedisPool redis.Pool

// 创建函数, 初始化Redis连接池
func InitRedis() {
	RedisPool = redis.Pool{
		MaxIdle:         20,
		MaxActive:       50,
		MaxConnLifetime: 60 * 5,
		IdleTimeout:     60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.81.128:6379")
		},
	}
}

func SaveImgCode(code, uuid string) error {
	// 1. 链接数据库
	conn, err := redis.Dial("tcp", "192.168.6.108:6379")
	if err != nil {
		fmt.Println("redis Dial err:", err)
		return err
	}
	defer conn.Close()

	// 2. 写数据库  --- 有效时间 5 分钟
	_, err = conn.Do("setex", uuid, 60*5, code)

	return err // 不需要回复助手!
}

// 校验图片验证码
func CheckImgCode(uuid, imgCode string) bool {
	// 链接 redis --- 从链接池中获取链接
	conn, err := redis.Dial("tcp", "192.168.6.108:6379")
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

// 存储短信验证码
//func SaveSmsCode(phone, code string) error {
//	// 链接 Redis --- 从链接池中获取一条链接
//	conn := RedisPool.Get()
//	defer conn.Close()
//
//	// 存储短信验证码到 redis 中
//	_, err := conn.Do("setex", phone+"_code", 60*3, code)
//
//	return err
//}

// 处理登录业务
func Login(mobile, pwd string) (string, error) {
	var user User
	err := GlobalConn.Where("mobile = ?", mobile).
		Where("password_hash = ?", pwd).
		Find(&user).Error

	return user.Name, err
}
