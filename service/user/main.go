package main

import (
	"go-micro.dev/v5"
	"go-micro.dev/v5/registry/consul"
	"user/handler"
	user "user/proto"
)

func main() {
	//初始化consul
	consulReg := consul.NewConsulRegistry()

	// Create service
	service := micro.NewService(
		micro.Name("micro_user"),
		micro.Registry(consulReg),             //添加注册
		micro.Address("192.168.81.128:12312"), //主动添加addr,防止其生成随机port
		micro.Version("latest"),
	)

	// 暂时用不上Initialize service
	//service.Init()
	// Register handler

	user.UserHandler(service.Server(), new(handler.User))

	// Run service
	service.Run()
}
