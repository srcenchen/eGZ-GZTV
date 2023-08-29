package v1

import "github.com/gogf/gf/v2/frame/g"

// GetVersionReq /** 获取版本信息
type GetVersionReq struct {
	g.Meta `method:"GET" summary:"获取版本信息" tags:"设置"`
}

// GetVersionRes /** 获取版本信息
type GetVersionRes struct {
	g.Meta  `method:"GET" summary:"获取版本信息" tags:"设置"`
	Version string `json:"version"`
}
