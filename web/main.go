package main

import (
	"gin_test01/web/controller"
	"gin_test01/web/model"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化Mysql连接池
	model.InitDb()
	//   1 初始化路由
	router := gin.Default()
	//   2 路由匹配
	// router.GET("/", func(c *gin.Context) {
	// 	c.Writer.WriteString("demo start")
	// })
	router.Static("/home", "web/view")
	// 此处GetSession是回调函数，只用写函数名
	router.GET("/api/v1.0/session", controller.GetSession)

	router.GET("/api/v1.0/imagecode/:uuid", controller.GetImageCd)
	//   3 启动运行
	router.Run(":8080")
}
