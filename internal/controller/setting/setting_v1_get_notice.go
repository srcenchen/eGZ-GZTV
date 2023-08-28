package setting

import (
	"context"
	"eGZ-GZTV/api/setting/v1"
	"eGZ-GZTV/internal/dao"
)

func (c *ControllerV1) GetNotice(_ context.Context, _ *v1.GetNoticeReq) (res *v1.GetNoticeRes, err error) {
	res = &v1.GetNoticeRes{
		Notice: dao.GetNotice(),
	}
	return
}
