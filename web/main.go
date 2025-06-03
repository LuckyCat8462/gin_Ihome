package main

import "github.com/gin-gonic/gin"

func main() {

	//   1 初始化路由
	router := gin.Default()
	//   2 路由匹配
	// router.GET("/", func(c *gin.Context) {
	// 	c.Writer.WriteString("demo start")
	// })
	router.Static("/", "web/view")

	//   3 启动运行
	router.Run(":8080")
}
