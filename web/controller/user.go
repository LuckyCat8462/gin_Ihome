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
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"image/png"
	"net/http"
)

//// 测试用-获取session信息
//func GetSession1(ctx *gin.Context) {
//	// 	初始化错误返回的map
//	resp := make(map[string]string)
//	// 调用utils包种的宏，go中称为常量
//	resp["errno"] = utils.RECODE_SESSIONERR
//	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
//
//	ctx.JSON(http.StatusOK, resp)
//}

func GetSession(ctx *gin.Context) {
	// 	初始化错误返回的map
	resp := make(map[string]interface{})
	// 调用utils包种的宏，go中称为常量
	s := sessions.Default(ctx)
	userName := s.Get("userName")
	if userName == nil {
		fmt.Println("用户未登录-getSession错了")
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		fmt.Println("session拿到了")
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		var nameData struct {
			Name string `json:"name"`
		}
		fmt.Println("nameData输出:", nameData)
		nameData.Name = userName.(string)
		resp["data"] = nameData
	}

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

// 获取验证码
func GetSmscd(ctx *gin.Context) {
	// 获取uuid
	phone := ctx.Param("phone")
	//	拆分GET请求中的url格式		资源路径?key1=v1&key2=v2
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")
	fmt.Println("out________", phone, imgCode, uuid)
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
	// 先从MySQL中获取数据.
	var areas []model.Area

	// 从缓存redis 中, 获取数据
	//conn := model.RedisPool.Get()
	conn, err := redis.Dial("tcp", "192.168.81.128:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 当初使用 "字节切片" 存入, 现在使用 切片类型接收
	areaData, _ := redis.Bytes(conn.Do("get", "areaData"))
	// 没有从 Redis 中获取到数据
	if len(areaData) == 0 {
		fmt.Println("从 MySQL 中 获取数据...")
		model.GlobalConn.Find(&areas)
		// 把数据写入到 redis 中. , 存储结构体序列化后的 json 串
		areaBuf, _ := json.Marshal(areas)
		conn.Do("set", "areaData", areaBuf)

	} else {
		fmt.Println("从 redis 中 获取数据...")
		// redis 中有数据
		json.Unmarshal(areaData, &areas)
	}

	resp := make(map[string]interface{})

	resp["errno"] = "0"
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas

	ctx.JSON(http.StatusOK, resp)
}

// 处理登录业务
func PostLogin(ctx *gin.Context) {
	//	获取前端数据
	var loginData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
	}
	ctx.Bind(&loginData)

	resp := make(map[string]interface{})

	//获取数据库数据,查询是否和数据库的数据匹配
	userName, err := model.Login(loginData.Mobile, loginData.PassWord)

	if err == nil {
		fmt.Println("postLogin登录成功")
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		s := sessions.Default(ctx) //初始化session
		//fmt.Println("初始化session")
		s.Set("userName", userName) //将用户名设置到session中
		s.Save()
	} else {
		fmt.Println("登录失败")
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
	}

	ctx.JSON(http.StatusOK, resp)
}

func DeleteSession(ctx *gin.Context) {
	resp := make(map[string]interface{})
	//初始化session对象
	s := sessions.Default(ctx)
	//删除session数据
	s.Delete("userName")
	//	必须使用save保存
	err := s.Save()
	if err != nil {
		fmt.Println("删除失败", err)
		resp["errno"] = utils.RECODE_IOERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_IOERR)
	} else {
		resp["errno"] = "0"
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	ctx.JSON(http.StatusOK, resp)
}
