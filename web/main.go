package main

import (
	"gin_test01/web/controller"
	"gin_test01/web/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化操作
	//#
	//初始化Mysql连接池
	model.InitDb()
	//初始化redis连接池
	model.InitRedis()
	//初始化容器
	store, _ := redis.NewStore(10, "tcp", "192.168.81.128:6379", "", "", []byte("secret"))
	//1.初始化路由
	router := gin.Default()
	//
	//
	//
	//2.路由匹配
	router.Static("/home", "web/view")
	//添加路由组
	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/session", controller.GetSession) // 此处GetSession是回调函数，只用写函数名
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:phone", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
		r1.GET("/areas", controller.GetArea)
		r1.POST("/sessions")
	}

	//使用容器
	router.Use(sessions.Sessions("sessionTest1", store))

	//3.启动运行
	router.Run(":8080")
}
