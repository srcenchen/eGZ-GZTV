package user

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/user/v1"
)

func (c *ControllerV1) LoginVerify(ctx context.Context, req *v1.LoginVerifyReq) (res *v1.LoginVerifyRes, err error) {
	isSuccess := dao.VerifyUser(req.Username, req.Password)
	res = &v1.LoginVerifyRes{
		IsSuccess: isSuccess,
	}
	return
}
