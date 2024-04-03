package setting

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/setting/v1"
)

func (c *ControllerV1) SetTitle(ctx context.Context, req *v1.SetTitleReq) (res *v1.SetTitleRes, err error) {
	dao.SetTitle(req.Title)
	return
}
