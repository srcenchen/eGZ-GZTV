package v1

import (
	"eGZ-GZTV/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// GetLiveListReq /** 获取直播列表
type GetLiveListReq struct {
	g.Meta `method:"GET" summary:"获取直播列表" tags:"直播"`
}

// GetLiveListRes /** 获取直播列表
type GetLiveListRes struct {
	g.Meta `method:"GET" summary:"获取直播列表" tags:"直播"`
	List   []entity.LiveTable `json:"list"` // 直播列表
}

// GetLiveReq /** 获取直播
type GetLiveReq struct {
	g.Meta `method:"GET" summary:"获取直播" tags:"直播"`
	Id     string `v:"required#直播ID不能为空"` // 直播ID
}

// GetLiveRes /** 获取直播
type GetLiveRes struct {
	g.Meta `method:"GET" summary:"获取直播" tags:"直播"`
	Live   entity.LiveTable `json:"live"` // 直播
}

// AddLiveReq /** 添加直播
type AddLiveReq struct {
	g.Meta      `method:"POST" summary:"添加直播" tags:"直播"`
	Title       string `v:"required#标题不能为空"` // 标题
	Description string `v:"required#描述不能为空"` // 描述
	HeadImage   string `v:"required#封面不能为空"` // 封面
}

// AddLiveRes /** 添加直播
type AddLiveRes struct {
	g.Meta `method:"POST" summary:"添加直播" tags:"直播"`
}

// DeleteLiveReq /** 删除直播
type DeleteLiveReq struct {
	g.Meta `method:"POST" summary:"删除直播" tags:"直播"`
	Id     string `v:"required#直播ID不能为空"` // 直播ID
}

// DeleteLiveRes /** 删除直播
type DeleteLiveRes struct {
	g.Meta `method:"POST" summary:"删除直播" tags:"直播"`
}
