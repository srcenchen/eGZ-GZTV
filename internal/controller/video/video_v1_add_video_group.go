package video

import (
	"context"
	"eGZ-GZTV/internal/dao"
	"time"

	"eGZ-GZTV/api/video/v1"
)

func (c *ControllerV1) AddVideoGroup(ctx context.Context, req *v1.AddVideoGroupReq) (res *v1.AddVideoGroupRes, err error) {
	t := time.Now()
	uploadDate := t.Format("2006-01-02 15:04:05")
	dao.AddVideo(req.Title, req.Description, "", "", uploadDate, -1)
	return
}
