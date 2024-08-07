package setting

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/setting/v1"
)

func (c *ControllerV1) GetFooter(ctx context.Context, req *v1.GetFooterReq) (res *v1.GetFooterRes, err error) {
	res = &v1.GetFooterRes{
		Footer: dao.GetFooter(),
	}
	return
}
