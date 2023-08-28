package live

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"eGZ-GZTV/api/live/v1"
)

func (c *ControllerV1) AddLive(ctx context.Context, req *v1.AddLiveReq) (res *v1.AddLiveRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
