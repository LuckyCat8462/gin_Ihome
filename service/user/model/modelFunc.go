package model

import (
	_ "github.com/go-sql-driver/mysql"
)

// 存储用户真实姓名
func SaveRealName(userName, realName, idCard string) error {
	return GlobalConn.Model(new(User)).Where("name = ?", userName).
		Updates(map[string]interface{}{"real_name": realName, "id_card": idCard}).Error
}

//func Login(mobile, pwd string) (string, error) {
//	var user User
//	err := GlobalConn.Where("mobile = ?", mobile).Select("name").
//		Where("password_hash = ?", pwd).
//		Find(&user).Error
//	//fmt.Println("测试用",user.Name)
//	return user.Name, err
//}
