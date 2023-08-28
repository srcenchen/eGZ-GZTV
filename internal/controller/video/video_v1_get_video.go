package video

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/video/v1"
)

func (c *ControllerV1) GetVideo(_ context.Context, req *v1.GetVideoReq) (res *v1.GetVideoRes, err error) {
	res = &v1.GetVideoRes{
		Video: dao.GetVideoByID(req.Id),
	}
	return
}
