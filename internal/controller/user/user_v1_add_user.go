package user

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/user/v1"
)

func (c *ControllerV1) AddUser(ctx context.Context, req *v1.AddUserReq) (res *v1.AddUserRes, err error) {
	dao.AddUser(req.Username, req.Password)
	return
}
