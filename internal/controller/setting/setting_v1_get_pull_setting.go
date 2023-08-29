package setting

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/setting/v1"
)

func (c *ControllerV1) GetPullSetting(_ context.Context, _ *v1.GetPullSettingReq) (res *v1.GetPullSettingRes, err error) {
	res = &v1.GetPullSettingRes{
		PullSetting: dao.GetPull(),
	}
	return
}
