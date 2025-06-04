package main

import (
	"gin_test01/web/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	//   1 初始化路由
	router := gin.Default()
	//   2 路由匹配
	// router.GET("/", func(c *gin.Context) {
	// 	c.Writer.WriteString("demo start")
	// })
	router.Static("/home", "web/view")
	// 此处GetSession是回调函数，只用写函数名
	router.GET("/api/v1.0/session", controller.GetSeesion)

	//   3 启动运行
	router.Run(":8080")
}
