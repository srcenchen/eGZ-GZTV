package setting

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/setting/v1"
)

func (c *ControllerV1) SetNotice(ctx context.Context, req *v1.SetNoticeReq) (res *v1.SetNoticeRes, err error) {
	dao.SetNotice(req.Notice)
	return
}
