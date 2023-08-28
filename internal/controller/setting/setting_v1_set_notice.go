package setting

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"eGZ-GZTV/api/setting/v1"
)

func (c *ControllerV1) SetNotice(ctx context.Context, req *v1.SetNoticeReq) (res *v1.SetNoticeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
