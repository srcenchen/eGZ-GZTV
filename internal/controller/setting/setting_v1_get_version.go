package setting

import (
	"context"
	"eGZ-GZTV/internal/consts"

	"eGZ-GZTV/api/setting/v1"
)

func (c *ControllerV1) GetVersion(ctx context.Context, req *v1.GetVersionReq) (res *v1.GetVersionRes, err error) {
	res = &v1.GetVersionRes{
		Version: consts.Version,
	}
	return
}
