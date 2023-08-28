package live

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"eGZ-GZTV/api/live/v1"
)

func (c *ControllerV1) GetLive(ctx context.Context, req *v1.GetLiveReq) (res *v1.GetLiveRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
