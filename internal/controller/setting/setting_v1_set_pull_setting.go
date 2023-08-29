package setting

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/setting/v1"
)

func (c *ControllerV1) SetPullSetting(_ context.Context, req *v1.SetPullSettingReq) (res *v1.SetPullSettingRes, err error) {
	dao.SetPull(req.PullSetting)
	return
}
