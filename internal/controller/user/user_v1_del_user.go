package user

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/user/v1"
)

func (c *ControllerV1) DelUser(ctx context.Context, req *v1.DelUserReq) (res *v1.DelUserRes, err error) {
	dao.DeleteUserByID(req.Id)
	return
}
