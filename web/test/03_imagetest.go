package main

import (
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"net/http"
)

func main() {
	// 初始化对象
	cap := captcha.New()
	// 设置字体
	cap.SetFont("web/conf/comic.ttf")
	// 	设置验证码的大小
	cap.SetSize(128, 64)
	// 	设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 	设置前景色
	cap.SetFrontColor(color.RGBA{128, 65, 0, 128})
	// 	设置背景颜色	可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	// 生成字体

	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		img, str := cap.Create(6, captcha.NUM)
		png.Encode(w, img)
		println(str)
	})

	http.ListenAndServe(":8086", nil)

}
