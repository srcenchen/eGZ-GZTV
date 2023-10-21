package video

import (
	"context"
	"eGZ-GZTV/internal/dao"

	"eGZ-GZTV/api/video/v1"
)

func (c *ControllerV1) GetVideoGroupList(ctx context.Context, req *v1.GetVideoGroupListReq) (res *v1.GetVideoGroupListRes, err error) {
	res = &v1.GetVideoGroupListRes{
		List: dao.GetVideoGroup(),
	}
	return
}
