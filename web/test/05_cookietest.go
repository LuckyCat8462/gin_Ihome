package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		//name 名称;
		//value 值;
		//maxAge：最大生命周期。
		//	 = 0 : 表示没指定该属性。
		//	 < 0 ：表示删除。 ---- // 删除Cookie 的操作， 可以使用 该属性实现。
		//	 > 0 ：指定生命周期。 单位：s
		//path：路径。—— 通常传""
		//domain：域名。 IP地址。
		//secure：设置是否安全保护。true：不能在 地址栏前，点击查看。 可以使用 F12 查看。
		//false：能在 地址栏前，点击查看。
		//httpOnly：是否只针对http协议。
		c.SetCookie("test1", "nekotest1", 60*60, "", "", true, true)
		c.Writer.WriteString("测试cookie...")
	})

	router.Run(":8089")
}
