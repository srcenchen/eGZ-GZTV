package video

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/video/v1"
)

func (c *ControllerV1) GetVideoList(_ context.Context, req *v1.GetVideoListReq) (res *v1.GetVideoListRes, err error) {
	res = &v1.GetVideoListRes{
		List: dao.GetVideos(),
	}
	return
}
