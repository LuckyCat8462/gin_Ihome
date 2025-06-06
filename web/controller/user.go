package controller

import (
	"fmt"
	"gin_test01/web/utils"
	"github.com/afocus/captcha"
	// ctrl左键点击captcha，可以看到其相关文件，包括examples、comic.ttf
	"github.com/gin-gonic/gin"
	"image/color"
	"image/png"
	"net/http"
)

// 获取session信息
func GetSession(ctx *gin.Context) {
	// 	初始化错误返回的map
	resp := make(map[string]string)
	// 调用utils包种的宏，go中称为常量
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSON(http.StatusOK, resp)
}

// 获取图片信息
func GetImageCd(ctx *gin.Context) {
	// 获取uuid
	uuid := ctx.Param("uuid")
	// 获取验证码
	// 初始化对象
	cap := captcha.New()
	// 设置字体
	cap.SetFont("web/conf/comic.ttf")
	// 	设置验证码的大小
	cap.SetSize(128, 64)
	// 	设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 	设置前景色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 	设置背景颜色	可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	// 生成字体
	img1, str1 := cap.Create(4, captcha.NUM) // 4位字
	png.Encode(ctx.Writer, img1)
	println(str1)
	fmt.Println("uuid", uuid)
}
