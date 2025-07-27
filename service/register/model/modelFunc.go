package model

import (
	_ "github.com/go-sql-driver/mysql"
)

// 注册用户信息,写 MySQL 数据库.
func RegisterUser(mobile, pwd string) error {
	var user User
	user.Mobile = mobile
	user.Name = mobile // 默认使用手机号作为用户名
	user.Password_hash = pwd

	// 插入数据到MySQL
	return GlobalConn.Create(&user).Error
}

// 处理登录业务
func Login(mobile, pwd string) (string, error) {
	var user User
	err := GlobalConn.Where("mobile = ?", mobile).Select("name").
		Where("password_hash = ?", pwd).
		Find(&user).Error
	//fmt.Println("测试用",user.Name)
	return user.Name, err
}
