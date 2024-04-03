package v1

import "github.com/gogf/gf/v2/frame/g"

// GetTitleReq /** 获取标题
type GetTitleReq struct {
	g.Meta `method:"GET" summary:"获取标题" tags:"设置"`
}

// GetTitleRes /** 获取标题
type GetTitleRes struct {
	g.Meta `method:"GET" summary:"获取标题" tags:"设置"`
	Title  string `json:"Title"` // 标题
}

// SetTitleReq /** 设置标题
type SetTitleReq struct {
	g.Meta `method:"GET" summary:"设置标题" tags:"设置"`
	Title  string // 标题
}

// SetTitleRes /** 设置标题
type SetTitleRes struct {
	g.Meta `method:"GET" summary:"设置标题" tags:"设置"`
}
