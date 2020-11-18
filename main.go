package main

import (
	"log"
	"os"

	"github.com/mrgAndysm/mic_user/protouser"
	"github.com/mrgAndysm/mic_user/service/grpc"

	"github.com/mrgAndysm/mic_user/apps/routers"
	"github.com/mrgAndysm/mic_user/com"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/web"
)

func main() {

	// consul 注册服务
	c := new(com.ConsulServer)
	c.ConsulServerInit()

	if os.Args[1] == "w" {
		//初始化路由
		ginRouter := routers.InitRouters()
		// web服务
		websrv := web.NewService(
			web.Handler(ginRouter),
			web.Name("userserver"),
			web.Address(":8001"),
			web.Handler(ginRouter),
			web.Registry(c.ConsulReg),
		)

		if err := websrv.Run(); err != nil {
			log.Fatal(err)
		}

	} else {
		// grpc微服务
		srv := micro.NewService(
			//micro.Name服务端向consul注册服务名称 客户端需要通过该名称来进行调用
			micro.Name("go.micro.svr.user"),
			micro.Version("latest"),
			micro.Registry(c.ConsulReg),
		)

		srv.Init()
		protouser.RegisterUserServiceHandler(srv.Server(), new(grpc.UserService))

		if err := srv.Run(); err != nil {
			log.Fatal(err)
		}

	}

}
