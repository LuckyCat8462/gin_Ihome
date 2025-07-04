package Learning

//package main
//
//import (
//  "github.com/gin-contrib/sessions"
//  "github.com/gin-contrib/sessions/redis"
//  "github.com/gin-gonic/gin"
//)
//
//func main() {
//  r := gin.Default()
//func NewStore(size int, network, address, password string, keyPairs ...[]byte) (Store, error)
//size:容器大小。
//network：协议
//address：IP：port
//password：使用redis做容器使用的密码。 没有特殊设定，传 “”
//[]byte(“secret”)： 加密密钥！
//  store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
//  r.Use(sessions.Sessions("mysession", store))
//
//  r.GET("/incr", func(c *gin.Context) {
//    session := sessions.Default(c)
//    var count int
//    v := session.Get("count")
//    if v == nil {
//      count = 0
//    } else {
//      count = v.(int)
//      count++
//    }
//    session.Set("count", count)
//    session.Save()
//    c.JSON(200, gin.H{"count": count})
//  })
//  r.Run(":8000")
//}
