package video

import (
	"context"
	"eGZ-GZTV/internal/dao"
	"eGZ-GZTV/internal/model"
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
	// 判断删除的是否是分组
	if video.GroupId == -1 {
		// 将分组下的视频移动到未分组
		// 获取分组下的视频
		videos := dao.GetVideoGroupByID(req.Id)
		// 遍历视频
		for _, v := range videos {
			// 修改视频的分组为未分组
			v.GroupId = -2
			// 保存视频
			model.GetDatabase().Save(&video)
		}
	}

	dao.DeleteVideoByID(req.Id)
	return
}
