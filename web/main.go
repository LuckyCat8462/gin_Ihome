package main

import (
	"gin_test01/web/controller"
	"gin_test01/web/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func LoginFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 初始化 Session 对象
		s := sessions.Default(ctx)
		userName := s.Get("userName")

		if userName == nil {
			ctx.Abort() // 从这里返回, 不必继续执行了
		} else {
			ctx.Next() // 继续向下
		}
	}
}

func main() {
	//初始化操作
	//#
	//初始化Mysql连接池
	model.InitDb()
	//初始化redis连接池
	model.InitRedis()
	//1.初始化路由
	router := gin.Default()

	//初始化容器
	//redis.NewStore(10, "tcp", "192.168.81.128:3479", "", "", []byte("secret"))
	store, _ := redis.NewStore(10, "tcp", "192.168.81.128:6379", "", "", []byte("secret"))
	//使用容器
	router.Use(sessions.Sessions("loginSession", store)) //使用中间件--指定容器

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
		r1.POST("/sessions", controller.PostLogin)

		r1.Use(LoginFilter()) //调用中间件--下方的路由都不需要校验session了,可以直接获取数据
		r1.DELETE("session", controller.DeleteSession)
		r1.GET("/user", controller.GetUserInfo)
		r1.PUT("/user/name", controller.PutUserInfo)
		r1.POST("user/avatar", controller.PostAvater)

		//	上传实名认证
		r1.POST("/user/auth", controller.PutUserAuth)
		// 获取实名认证
		r1.GET("/user/auth", controller.GetUserInfo)

	}

	//3.启动运行
	router.Run(":8080")
}
