package main

import (
	"fmt"
	"gorm.io/gorm/schema"
	"time"

	//_ "github.com/go-sql-driver/mysql"
	//"_"表示不直接使用包，底层连接要使用
	_ "gin_test01/web/test2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 3.创建全局结构体
type Student struct {
	gorm.Model //go语言中匿名成员--继承
	Id         int
	Name       string `gorm:"size:50;default:'tester1'"` //string 默认大小是255,在创建时，可以指定为其他的大小
	Age        int
	Join       time.Time `gorm:"type:datetime"`
}

//- 默认MySQL数据库 有 3 种时间：
//  - date：
//  - datetime：
//  - timeStamp：时间戳。 ——  gorm 中，只有 timeStamp
//
//- 如果必须使用 MySQL 数据库特有的 “数据类型”， 使用 “type” 关键字来设置。

// 4.连接池-1-创建全局的连接池句柄
var GlobalConn *gorm.DB

func main() {
	fmt.Println("this is mian")

	//	1.连接数据库	格式：	用户名:密码@协议（ip：port）/数据库名
	//conn, err := gorm.Open("mysql", "neko:neko123456@tcp(127.0.0.1:3306/test)")
	dsn := "neko:neko123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ //	gorm参数
		//不要复数表名
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("连接数据库成功。。。")

	GlobalConn = db
	//2.借助gorm创建数据库表
	sqlDB, err := GlobalConn.DB()
	if err != nil {
		fmt.Println("获取数据库对象失败")
		return
	}
	//连接池-2
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	////创建表
	GlobalConn.AutoMigrate(&Student{})
	////创建数据
	//Stu := Student{Name: "neko", Age: 20}
	//result := GlobalConn.Create(&Stu) //通过指针创建

	//返回插入数据的name
	//name := Stu.Name
	//返回插入记录的条数
	//affect := result.RowsAffected
	//fmt.Println(name, affect)
	//ct := db.AutoMigrate(new(Student)).Error()
	//fmt.Println(ct)

	//增删改查
	InsertData()
}

// gorm增加数据
func InsertData() {
	Stu := Student{Name: "zhi", Age: 20}
	result := GlobalConn.Create(&Stu) //通过指针创建

	//返回插入数据的name
	fmt.Println(result.RowsAffected)
}

// gorm查询数据
func SelectData() {
	//获取第一条记录（主键升序）
	//SELECT * FROM users ORDER BY id LIMIT 1;
	GlobalConn.First(&Student{})

	//获取一条记录，没有指定排序字段
	// SELECT * FROM users LIMIT 1;
	GlobalConn.Take(&Student{})

	// 获取最后一条记录（主键降序）
	//SELECT * FROM users ORDER BY id DESC LIMIT 1;
	GlobalConn.Last(&Student{})

	//获取所有记录
	//需要将stu改为切片
	var stu1 []Student
	GlobalConn.Select("name", "age").Find(&stu1)
	fmt.Println(stu1)

	//where简易sql
	GlobalConn.Where("name = ?", "zhi").First(&stu1)
}

func UpdateData() {
	var Stu Student

	Stu.Name = "lisi4"
	Stu.Age = 22
	fmt.Println(GlobalConn.Model(&Student{}).
		Where("name = ?", "lisi4").
		Updates("lisi5").Error)
}

//	### 删除
//
// - 物理删除。真正的执行 Delete。
//
// ### 软删除
// - 逻辑删除。不真正删。不执行Delete。 —— 数据无价的！
//   - 创建表时，在表中添加一个 “删除字段” 。当需要删除时，更新 “删除字段”， 更新为 ：true
//   - 查询时，不查询 “删除字段” 为 null 的值。
//   - 举例：淘宝，删除订单。 手机，删除图片。
//
// - 实现 软删除：
func DeleteData() {
	GlobalConn.Where("name = ?", "lisi4").Delete(&Student{})

}
