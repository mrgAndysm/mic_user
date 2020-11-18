# 基于go-micro v2.0 实现一个用户微服务
    github.com/go-micro/micro/v2

### 注意： 
    基于v2版本，注意v3版本默认不支持consul

### 用于不同模块仓库互相引用 代码中引入了 github.com/mrgAndysm/mic_order 这个仓库的protoc
代码示例：

    package protocode
    
    import (
    	micro "github.com/micro/go-micro/v2"
    	registry "github.com/micro/go-micro/v2/registry"
    	consul "github.com/micro/go-plugins/registry/consul/v2"
    	protoorder "github.com/mrgAndysm/mic_order/protoorder"
    )
    
    // OrderServer order服务
    type OrderServer struct {
    	OrderClient protoorder.OrderService
    }
    
    // OrderServerInit 初始化order服务
    func (s *OrderServer) OrderServerInit() {
    
    	//新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口s
    	Reg := consul.NewRegistry(func(op *registry.Options) {
    		op.Addrs = []string{
    			"0.0.0.0:8500",
    		}
    	})
    
    	service := micro.NewService(
    		//micro.Name服务端向consul注册服务名称 客户端需要通过该名称来进行调用
    		micro.Name("go.micro.svr.order"),
    		micro.Version("latest"),
    		micro.Registry(Reg),
    	)
    
    	service.Init()
    
    	s.OrderClient = protoorder.NewOrderService("go.micro.svr.order", service.Client())
    
    }


