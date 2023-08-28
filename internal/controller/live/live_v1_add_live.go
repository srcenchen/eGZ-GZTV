package live

import (
	"context"
	"eGZ-GZTV/internal/dao"
	"strconv"
	"time"

	"eGZ-GZTV/api/live/v1"
)

func (c *ControllerV1) AddLive(_ context.Context, req *v1.AddLiveReq) (res *v1.AddLiveRes, err error) {
	timeUnix := time.Now().UnixNano()
	t := time.Now()
	uploadDate := t.Format("2006-01-02 15:04:05")
	dao.AddLive(req.Title, "liveGZTV", strconv.FormatInt(timeUnix, 10), req.HeadImage, req.Description, uploadDate)
	return
}
