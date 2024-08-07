package setting

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/setting/v1"
)

func (c *ControllerV1) SetFooter(ctx context.Context, req *v1.SetFooterReq) (res *v1.SetFooterRes, err error) {
	dao.SetFooter(req.Footer)
	return
}
