package video

import (
	"context"
	"eGZ-GZTV/internal/dao"
	"eGZ-GZTV/internal/model"

	"eGZ-GZTV/api/video/v1"
)

func (c *ControllerV1) EditVideoDetail(ctx context.Context, req *v1.EditVideoDetailReq) (res *v1.EditVideoDetailRes, err error) {
	var video = dao.GetVideoByID(req.Id)
	video.Title = req.Title
	video.GroupId = req.GroupID
	video.Description = req.Description
	model.GetDatabase().Save(&video)
	return
}
