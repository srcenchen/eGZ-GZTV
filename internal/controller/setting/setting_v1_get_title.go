package setting

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/setting/v1"
)

func (c *ControllerV1) GetTitle(ctx context.Context, req *v1.GetTitleReq) (res *v1.GetTitleRes, err error) {
	res = &v1.GetTitleRes{
		Title: dao.GetTitle(),
	}
	return
}
