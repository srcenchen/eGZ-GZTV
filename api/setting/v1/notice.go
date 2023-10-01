package v1

import "github.com/gogf/gf/v2/frame/g"

// GetNoticeReq /** 获取公告
type GetNoticeReq struct {
	g.Meta `method:"GET" summary:"获取公告" tags:"设置"`
}

// GetNoticeRes /** 获取公告
type GetNoticeRes struct {
	g.Meta `method:"GET" summary:"获取公告" tags:"设置"`
	Notice string `json:"notice"` // 公告
}

// SetNoticeReq /** 设置公告
type SetNoticeReq struct {
	g.Meta `method:"GET" summary:"设置公告" tags:"设置"`
	Notice string // 公告
}

// SetNoticeRes /** 设置公告
type SetNoticeRes struct {
	g.Meta `method:"GET" summary:"设置公告" tags:"设置"`
}
