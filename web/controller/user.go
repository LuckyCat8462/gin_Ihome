package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"gin_test01/web/model"
	getCaptcha "gin_test01/web/proto/getCaptcha"

	registerMicro "gin_test01/web/proto/register"
	"gin_test01/web/utils"
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/tedcy/fdfs_client"
	"image/png"
	"net/http"
	"path"
)

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
		fmt.Println("---GetSession拿到了---")
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		var nameData struct {
			Name string `json:"name"`
		}
		nameData.Name = userName.(string)
		resp["data"] = nameData
		fmt.Println("用户名：", nameData)
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
func PostReg(ctx *gin.Context) {
	// 获取数据
	var regData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
	}

	ctx.Bind(&regData)

	fmt.Println("测试reg数据：", regData.Mobile, regData.PassWord)

	//调用远程服务
	// 初始化consul
	microService := utils.InitMicro()
	microClient := registerMicro.NewRegisterService("micro_register", microService.Client())

	// 调用远程函数
	resp, err := microClient.Register(context.TODO(), &registerMicro.RegRequest{
		Mobile:   regData.Mobile,
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
	fmt.Println("user.go-GetArea函数：获取地址信息")
	// 没有从 Redis 中获取到数据
	if len(areaData) == 0 {
		fmt.Println("从 MySQL 中 获取数据...")
		model.GlobalConn.Find(&areas)
		// 把数据写入到 redis 中. , 存储结构体序列化后的 json 串
		areaBuf, _ := json.Marshal(areas)
		conn.Do("set", "areaData", areaBuf)

	} else {
		fmt.Println("从 Redis 中 获取数据...")
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

	//微服务版本登录功能
	////	获取前端数据
	var loginData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
	}
	ctx.Bind(&loginData)
	//fmt.Println("测试loginData", loginData)

	//指定consul服务发现
	microService := utils.InitMicro()
	//初始化客户端
	microClient := registerMicro.NewRegisterService("micro_register", microService.Client())

	resp, err := microClient.Login(context.TODO(), &registerMicro.RegRequest{Mobile: loginData.Mobile, Password: loginData.PassWord})
	fmt.Println("resp info", resp)
	if err != nil {
		fmt.Println("调用login服务错误", err)
		return
	} else {
		fmt.Println("login服务正常")
		//返回数据  存储session  并返回数据给web端
		session := sessions.Default(ctx)
		session.Set("userName", resp.Name)
		session.Save()
		//fmt.Println("login测试session", session.Get("userName"))
	}

	ctx.JSON(http.StatusOK, resp)

	//非微服务版本
	//resp := make(map[string]interface{})
	//
	////获取数据库数据,查询是否和数据库的数据匹配
	//userName, err := model.Login(loginData.Mobile, loginData.PassWord)
	//
	//if err == nil {
	//	fmt.Println("postLogin登录成功")
	//	resp["errno"] = utils.RECODE_OK
	//	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	//
	//	s := sessions.Default(ctx) //初始化session
	//	//fmt.Println("初始化session")
	//	s.Set("userName", userName) //将用户名设置到session中
	//	s.Save()
	//} else {
	//	fmt.Println("登录失败")
	//	resp["errno"] = utils.RECODE_LOGINERR
	//	resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
	//}
	//
	//ctx.JSON(http.StatusOK, resp)
}

// 退出登录
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

// 获取用户信息
func GetUserInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	// 获取 Session, 得到 当前 用户信息
	s := sessions.Default(ctx)
	userName := s.Get("userName")
	// 判断用户名是否存在.
	if userName == nil { // 用户没登录, 但进入该页面, 恶意进入.
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return // 如果出错, 报错, 退出
	}

	// 根据用户名, 获取 用户信息  ---- 查 MySQL 数据库  user 表.
	user, err := model.GetUserInfo(userName.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return // 如果出错, 报错, 退出
	}

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	temp := make(map[string]interface{})
	temp["user_id"] = user.ID
	temp["name"] = user.Name
	temp["mobile"] = user.Mobile
	temp["real_name"] = user.Real_name
	temp["id_card"] = user.Id_card
	temp["avatar_url"] = "http://192.168.81.128:8089" + user.Avatar_url

	resp["data"] = temp
}

// 更新用户名
func PutUserInfo(ctx *gin.Context) {
	// 获取当前用户名
	s := sessions.Default(ctx) // 初始化Session 对象
	userName := s.Get("userName")

	// 获取新用户名		---- 处理 Request Payload 类型数据. Bind()
	var nameData struct {
		Name string `json:"name"`
	}
	ctx.Bind(&nameData)

	// 更新用户名
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	// 更新数据库中的 name
	err := model.UpdateUserName(nameData.Name, userName.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}

	// 更新 Session 数据
	s.Set("userName", nameData.Name)
	err = s.Save() // 必须保存
	if err != nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = nameData
}

// 上传头像
func PostAvater(ctx *gin.Context) {
	////	获取图片文件,静态文件
	file, _ := ctx.FormFile("avatar")
	////	上传文件到项目中
	//err := ctx.SaveUploadedFile(FileHeader, "test/"+FileHeader.Filename)
	//fmt.Println(err)

	//上传文件到fdfs中
	clt, _ := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")
	//打开文件，读取文件内容
	f, _ := file.Open()

	buf := make([]byte, file.Size)
	f.Read(buf)
	//go语言根据文件名获取文件后缀
	fileExt := path.Ext(file.Filename) //传文件名
	//按字节流上传图片内容
	remoteId, _ := clt.UploadByBuffer(buf, fileExt[1:]) //fileExt是一个切片，所以可以从下标1取

	//获取session，得到当前用户
	userName := sessions.Default(ctx).Get("userName")

	//根据用户名,更新用户头像
	model.UpdateAvatar(userName.(string), remoteId)

	resp := make(map[string]interface{})
	resp["errno"] = "0"
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	temp := make(map[string]interface{})
	temp["avatar_url"] = "http://192.168.81.128:8089" + remoteId

	resp["data"] = temp
	ctx.JSON(http.StatusOK, resp)
}

type AuthStu struct {
	IdCard   string `json:"id_card"`
	RealName string `json:"real_name"`
}

// 上传实名认证
func PutUserAuth(ctx *gin.Context) {
	//获取前端数据
	var auth AuthStu
	err := ctx.Bind(&auth)

	//校验数据
	if err != nil {
		fmt.Println("获取数据错误", err)
		return
	}

	//session := sessions.Default(ctx)
	//userName := session.Get("userName")

	////处理数据 微服务
	//microService := utils.InitMicro()
	//microClient := userMicro.NewUserService("micro_user", microService.Client())
	//
	////调用远程服务
	//
	//resp, err := microClient.AuthUpdate(context.TODO(), &userMicro.AuthReq{
	//	UserName: userName.(string),
	//	RealName: auth.RealName,
	//	IdCard:   auth.IdCard,
	//})
	//if err != nil {
	//	fmt.Println("resp情况", err)
	//}

	//返回数据
	//ctx.JSON(http.StatusOK, resp)
}
