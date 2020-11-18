package routers

import (
	"fmt"
	"strconv"

	"github.com/mrgAndysm/mic_order/protoorder"
	"github.com/mrgAndysm/mic_user/com/protocode"

	"github.com/gin-gonic/gin"
)

// InitRouters 初始化路由
func InitRouters() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.GET("/order", func(context *gin.Context) {

		order := new(protocode.OrderServer)
		order.OrderServerInit()

		res, err := order.OrderClient.CreateOrder(context.Request.Context(), &protoorder.QOrderInfo{
			Price:  1,
			GoodId: 5,
		})
		if err != nil {
			fmt.Println(err)
		}
		context.String(200, "get orderinfos"+strconv.FormatInt(res.OrderId, 10))
	})

	return ginRouter
}
