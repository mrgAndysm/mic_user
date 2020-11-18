package com

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

// ConsulServer consul服务
type ConsulServer struct {
	ConsulReg registry.Registry
}

// ConsulServerInit 初始化consul服
func (s *ConsulServer) ConsulServerInit() {

	//新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口s
	s.ConsulReg = consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"0.0.0.0:8500",
		}
	})
}
