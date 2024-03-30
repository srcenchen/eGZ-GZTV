package video

import (
	"context"
	"eGZ-GZTV/api/video/v1"
	"eGZ-GZTV/internal/dao"
)

func (c *ControllerV1) AddVideoGroup(ctx context.Context, req *v1.AddVideoGroupReq) (res *v1.AddVideoGroupRes, err error) {
	dao.AddGroup(req.Title, req.Description, req.ParentGroup)
	return
}
