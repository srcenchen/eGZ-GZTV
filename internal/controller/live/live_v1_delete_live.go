package live

import (
	"context"
	"eGZ-GZTV/internal/dao"
	"os"

	"eGZ-GZTV/api/live/v1"
)

func (c *ControllerV1) DeleteLive(_ context.Context, req *v1.DeleteLiveReq) (res *v1.DeleteLiveRes, err error) {
	// 获取信息
	live := dao.GetLiveByID(req.Id)
	// 删除封面
	_ = os.Remove("./resource/upload/images/" + live.HeadImage)
	dao.DeleteLiveByID(req.Id)
	return
}
