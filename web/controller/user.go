package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"gin_test01/web/model"
	getCaptcha "gin_test01/web/proto/getCaptcha"
	userMicro "gin_test01/web/proto/user"
	"gin_test01/web/utils"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
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
	microService := utils.InitMicro()
	//初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("getCaptcha", microService.Client())
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

// 发送注册信息
func PostRet(ctx *gin.Context) {
	// 获取数据
	var regData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
		//SmsCode  string `json:"sms_code"`
	}

	ctx.Bind(&regData)

	//调用远程服务
	// 初始化consul
	microService := utils.InitMicro()
	microClient := userMicro.NewUserService("micro_user", microService.Client())

	// 调用远程函数
	resp, err := microClient.Register(context.TODO(), &userMicro.RegReq{
		Mobile: regData.Mobile,
		//SmsCode:  regData.SmsCode,
		Password: regData.PassWord,
	})
	if err != nil {
		fmt.Println("注册用户, 找不到远程服务!", err)
		return
	}
	// 写给浏览器
	ctx.JSON(http.StatusOK, resp)
}

// 获取地域信息
func GetArea(ctx *gin.Context) {
	//先从mysql中获取数据
	var areas []model.Area
	result := model.GlobalConn.Find(&areas)
	fmt.Println(result.RowsAffected)
	//	把数据写到redis中
	//conn := model.RedisPool.Get() //获取链接
	conn, err := redis.Dial("tcp", "192.168.81.128:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Do("set", "areaData", areas)

	resp := make(map[string]interface{})
	resp["errno"] = "0"
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas

	ctx.JSON(http.StatusOK, resp)
}

//func GetArea(ctx *gin.Context) {
//	// 先从MySQL中获取数据.
//	var areas []model.Area
//
//	// 从缓存redis 中, 获取数据
//	conn := model.RedisPool.Get()
//	// 当初使用 "字节切片" 存入, 现在使用 切片类型接收
//	areaData, _ := redis.Bytes(conn.Do("get", "areaData"))
//	// 没有从 Redis 中获取到数据
//	if len(areaData) == 0 {
//
//		fmt.Println("从 MySQL 中 获取数据...")
//		model.GlobalConn.Find(&areas)
//		// 把数据写入到 redis 中. , 存储结构体序列化后的 json 串
//		areaBuf, _ := json.Marshal(areas)
//		conn.Do("set", "areaData", areaBuf)
//
//	} else {
//		fmt.Println("从 redis 中 获取数据...")
//		// redis 中有数据
//		json.Unmarshal(areaData, &areas)
//	}
//
//	resp := make(map[string]interface{})
//
//	resp["errno"] = "0"
//	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
//	resp["data"] = areas
//
//	ctx.JSON(http.StatusOK, resp)
//}
