package user

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/user/v1"
)

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	res = &v1.UserListRes{
		List: dao.GetUserList(),
	}
	return
}
