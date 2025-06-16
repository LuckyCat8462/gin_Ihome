package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"
	"go-micro.dev/v5"
	"go-micro.dev/v5/registry/consul"
)

func main() {
	//初始化consul
	consulReg := consul.NewConsulRegistry()

	// Create service
	service := micro.NewService(
		micro.Name("getcaptcha"),
		micro.Registry(consulReg),             //添加注册
		micro.Address("192.168.81.128:18546"), //主动添加addr,防止其生成随机port
		micro.Version("latest"),
	)

	// 暂时用不上Initialize service
	//service.Init()
	// Register handler

	pb.RegisterGetCaptchaHandler(service.Server(), handler.New())

	// Run service
	service.Run()
}
