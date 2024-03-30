package video

import (
	"context"
	"eGZ-GZTV/api/video/v1"
	"eGZ-GZTV/internal/dao"
	"eGZ-GZTV/internal/model"
	"strconv"
)

func (c *ControllerV1) DelVideoGroup(ctx context.Context, req *v1.DelVideoGroupReq) (res *v1.DelVideoGroupRes, err error) {
	DelGroup(req.Id)
	return
}

func DelGroup(groupId string) {
	dao.DeleteGroupByID(groupId)
	// 获取分组详情
	videos := dao.GetVideoByGroupID(groupId)
	// 遍历视频
	for _, v := range videos {
		// 修改视频的分组为未分组
		v.GroupId = -2
		// 保存视频
		model.GetDatabase().Save(&v)
	}
	// 获取依赖于他的子分类
	groups := dao.GetGroupByParent(groupId)
	for _, v := range groups {
		DelGroup(strconv.Itoa(v.Id))
	}

}
