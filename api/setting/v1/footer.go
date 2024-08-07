package v1

import "github.com/gogf/gf/v2/frame/g"

// GetFooterReq /** 获取页脚
type GetFooterReq struct {
	g.Meta `method:"GET" summary:"获取页脚" tags:"设置"`
}

// GetFooterRes /** 获取页脚
type GetFooterRes struct {
	g.Meta `method:"GET" summary:"获取页脚" tags:"设置"`
	Footer string `json:"Footer"` // 页脚
}

// SetFooterReq /** 设置页脚
type SetFooterReq struct {
	g.Meta `method:"GET" summary:"设置页脚" tags:"设置"`
	Footer string // 页脚
}

// SetFooterRes /** 设置页脚
type SetFooterRes struct {
	g.Meta `method:"GET" summary:"设置页脚" tags:"设置"`
}
