package controller

import (
	"context"
	"encoding/json"
	"fmt"
	getCaptcha "gin_test01/web/proto"
	"gin_test01/web/utils"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v5"
	"go-micro.dev/v5/registry/consul"
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

	//指定consul服务发现
	consulReg := consul.NewConsulRegistry()
	consulService := micro.NewService(micro.Registry(consulReg))
	//初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("getCaptcha", consulService.Client())
	//调用远程服务
	//microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		fmt.Println("未找到远程服务。。。")
		fmt.Println(err)
		return
	}
	//resp并不能直接使用，因为传入服务端时进行了序列化，此时还需要将得到的数据进行反序列化
	var img captcha.Image
	//img := captcha.Image{}
	json.Unmarshal(resp.Img, &img)
	//将图片写入到浏览器

	png.Encode(ctx.Writer, img)
	//fmt.Println("str", string(resp.Img))
	fmt.Println("uuid", uuid)
}
