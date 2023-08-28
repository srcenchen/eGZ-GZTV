package setting

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"eGZ-GZTV/api/setting/v1"
)

func (c *ControllerV1) GetNotice(ctx context.Context, req *v1.GetNoticeReq) (res *v1.GetNoticeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
