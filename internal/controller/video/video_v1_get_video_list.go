package video

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"eGZ-GZTV/api/video/v1"
)

func (c *ControllerV1) GetVideoList(ctx context.Context, req *v1.GetVideoListReq) (res *v1.GetVideoListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
