package grpc

import (
	"context"

	"github.com/mrgAndysm/mic_user/protouser"
)

// UserService 订单服务
type UserService struct{}

// CreateUser 创建
func (o *UserService) CreateUser(ctx context.Context, req *protouser.QUserInfo, rsp *protouser.RUserResult) error {
	//  req.Phone
	// req.UserName
	rsp.UserId = 1
	return nil
}
