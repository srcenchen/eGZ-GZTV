package live

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"eGZ-GZTV/api/live/v1"
)

func (c *ControllerV1) GetLiveList(ctx context.Context, req *v1.GetLiveListReq) (res *v1.GetLiveListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
