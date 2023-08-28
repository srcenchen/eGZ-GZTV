package user

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/user/v1"
)

func (c *ControllerV1) ChangePassword(_ context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error) {
	dao.ChangePassword(req.Username, req.NewPassword)
	return
}
