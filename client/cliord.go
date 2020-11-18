package main

import (
	"context"
	"fmt"

	"github.com/mrgAndysm/mic_user/com"
	"github.com/mrgAndysm/mic_user/protouser"

	"github.com/micro/go-micro/v2"
)

func main() {
	//新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
	c := new(com.ConsulServer)
	c.ConsulServerInit()

	service := micro.NewService(
		//micro.Name服务端向consul注册服务名称 客户端需要通过该名称来进行调用
		micro.Name("go.micro.svr.user"),
		micro.Version("latest"),
		micro.Registry(c.ConsulReg),
	)

	service.Init()

	user := protouser.NewUserService("go.micro.svr.user", service.Client())

	res, err := user.CreateUser(context.TODO(), &protouser.QUserInfo{
		UserName: "PGMAN",
		Phone:    "13538462645",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.UserId)
}
