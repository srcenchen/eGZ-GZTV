package v1

import "github.com/gogf/gf/v2/frame/g"

// GetPullSettingReq /** 获取拉流设置
type GetPullSettingReq struct {
	g.Meta `method:"GET" summary:"获取拉流设置" tags:"设置"`
}

// GetPullSettingRes /** 获取拉流设置
type GetPullSettingRes struct {
	g.Meta      `method:"GET" summary:"获取拉流设置" tags:"设置"`
	PullSetting string `json:"pull_setting"` // 公告
}

// SetPullSettingReq /** 设置拉流设置
type SetPullSettingReq struct {
	g.Meta      `method:"GET" summary:"设置拉流设置" tags:"设置"`
	PullSetting string `v:"required#拉流设置不能为空" json:"pull_setting"`
}

// SetPullSettingRes /** 设置拉流设置
type SetPullSettingRes struct {
	g.Meta `method:"GET" summary:"设置拉流设置" tags:"设置"`
}
