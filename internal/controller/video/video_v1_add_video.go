package video

import (
	"context"
	"eGZ-GZTV/internal/dao"
	"time"

	"eGZ-GZTV/api/video/v1"
)

func (c *ControllerV1) AddVideo(_ context.Context, req *v1.AddVideoReq) (res *v1.AddVideoRes, err error) {
	t := time.Now()
	uploadDate := t.Format("2006-01-02 15:04:05")
	if req.GroupId == 0 {
		req.GroupId = -2
	}
	dao.AddVideo(req.Title, req.Description, req.VideoName, req.HeadImage, uploadDate, req.GroupId)
	return
}
