package v1

import (
	"eGZ-GZTV/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// GetVideoListReq /** 获取视频列表
type GetVideoListReq struct {
	g.Meta `method:"GET" summary:"获取视频列表" tags:"视频"`
}

// GetVideoListRes /** 获取视频列表
type GetVideoListRes struct {
	g.Meta `method:"GET" summary:"获取视频列表" tags:"视频"`
	List   []entity.VideoTable `json:"list"` // 视频列表
}

// GetVideoReq /** 获取视频
type GetVideoReq struct {
	g.Meta `method:"GET" summary:"获取视频" tags:"视频"`
	Id     string `v:"required#视频ID不能为空"` // 视频ID
}

// GetVideoRes /** 获取视频
type GetVideoRes struct {
	g.Meta `method:"GET" summary:"获取视频" tags:"视频"`
	Video  entity.VideoTable `json:"video"` // 视频
}

// AddVideoReq /** 添加视频
type AddVideoReq struct {
	g.Meta      `method:"POST" summary:"添加视频" tags:"视频"`
	Title       string `v:"required#标题不能为空"` // 标题
	Description string `v:"required#描述不能为空"` // 描述
	VideoName   string `v:"required#视频不能为空"` // 视频名称
	HeadImage   string `v:"required#封面不能为空"` // 封面
	GroupId     int    `v:""`                // 视频组
}

// AddVideoRes /** 添加视频
type AddVideoRes struct {
	g.Meta `method:"POST" summary:"添加视频" tags:"视频"`
}

// DeleteVideoReq /** 删除视频
type DeleteVideoReq struct {
	g.Meta `method:"POST" summary:"删除视频" tags:"视频"`
	Id     string `v:"required#视频ID不能为空"` // 视频ID
}

// DeleteVideoRes /** 删除视频
type DeleteVideoRes struct {
	g.Meta `method:"POST" summary:"删除视频" tags:"视频"`
}

// AddVideoGroupReq /** 添加视频组
type AddVideoGroupReq struct {
	g.Meta      `method:"POST" summary:"添加视频组" tags:"视频"`
	Title       string `v:"required#标题不能为空"` // 标题
	Description string `v:"required#描述不能为空"` // 描述
}

// AddVideoGroupRes /** 添加视频组
type AddVideoGroupRes struct {
	g.Meta `method:"POST" summary:"添加视频组" tags:"视频"`
}

// GetVideoGroupListReq /** 获取视频组列表
type GetVideoGroupListReq struct {
	g.Meta `method:"GET" summary:"获取视频组列表" tags:"视频"`
}

// GetVideoGroupListRes /** 获取视频组列表
type GetVideoGroupListRes struct {
	g.Meta `method:"GET" summary:"获取视频组列表" tags:"视频"`
	List   []entity.VideoTable `json:"list"` // 视频列表
}

// GetVideoGroupReq /** 获取视频组
type GetVideoGroupReq struct {
	g.Meta `method:"GET" summary:"获取视频组" tags:"视频"`
	Id     string `v:"required#视频组ID不能为空"` // 视频组ID
}

// GetVideoGroupRes /** 获取视频组
type GetVideoGroupRes struct {
	g.Meta `method:"GET" summary:"获取视频组" tags:"视频"`
	List   []entity.VideoTable `json:"list"` // 视频列表
}

// EditVideoDetailReq /** 修改视频信息
type EditVideoDetailReq struct {
	g.Meta      `method:"POST" summary:"修改视频信息" tags:"视频"`
	Id          string `v:"required#视频ID不能为空"`    // 视频ID
	Title       string `v:"required#标题不能为空"`      // 标题
	Description string `v:"required#描述不能为空"`      // 描述
	GroupID     int    `v:"required#GroupID不能为空"` // GroupID
}

// EditVideoDetailRes /** 修改视频信息
type EditVideoDetailRes struct {
	g.Meta `method:"POST" summary:"修改视频信息" tags:"视频"`
}
