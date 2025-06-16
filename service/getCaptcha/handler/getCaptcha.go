package handler

import (
	"context"
	"encoding/json"
	pb "getCaptcha/proto"
	"github.com/afocus/captcha"
	"image/color"
)

type GetCaptcha struct{}

// Return a new handler
func New() *GetCaptcha {
	return &GetCaptcha{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetCaptcha) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	// 生成图片验证码
	// 初始化对象

	cap := captcha.New()
	// 设置字体
	cap.SetFont("./conf/comic.ttf")
	// 	设置验证码的大小
	cap.SetSize(128, 64)
	// 	设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 	设置前景色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 	设置背景颜色	可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	// 生成字体
	img1, _ := cap.Create(4, captcha.NUM) // 4位字
	//将生成的图片进行序列化为json文件，用于传输
	imgBuf, _ := json.Marshal(img1)

	//将imgBuf 使用参数rsp 传出

	rsp.Img = imgBuf

	return nil
}
