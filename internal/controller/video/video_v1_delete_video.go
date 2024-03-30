package video

import (
	"context"
	"eGZ-GZTV/internal/dao"
	"os"

	"eGZ-GZTV/api/video/v1"
)

func (c *ControllerV1) DeleteVideo(_ context.Context, req *v1.DeleteVideoReq) (res *v1.DeleteVideoRes, err error) {
	// 获取视频信息
	video := dao.GetVideoByID(req.Id)
	// 删除视频
	_ = os.Remove("./resource/upload/videos/" + video.VideoName)
	// 删除封面
	_ = os.Remove("./resource/upload/images/" + video.HeadImage)

	dao.DeleteVideoByID(req.Id)
	return
}
