package video

import (
	"context"
	"eGZ-GZTV/api/video/v1"
)

func (c *ControllerV1) GetVideoGroup(ctx context.Context, req *v1.GetVideoGroupReq) (res *v1.GetVideoGroupRes, err error) {
	//res = &v1.GetVideoGroupRes{
	//	List: dao.GetVideoByGroupID(req.Id),
	//}
	return
}
