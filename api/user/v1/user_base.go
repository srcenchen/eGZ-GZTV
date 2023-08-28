package v1

import "github.com/gogf/gf/v2/frame/g"

// LoginVerifyReq /** 登录验证
type LoginVerifyReq struct {
	g.Meta   `method:"POST" summary:"登录验证" tags:"用户"`
	Username string `v:"required#用户名不能为空"` // 用户名
	Password string `v:"required#密码不能为空"`  // 密码
}

// LoginVerifyRes /** 登录验证
type LoginVerifyRes struct {
	g.Meta    `method:"POST" summary:"登录验证" tags:"用户"`
	IsSuccess bool `json:"isSuccess"` // token
}

// ChangePasswordReq /** 修改密码
type ChangePasswordReq struct {
	g.Meta      `method:"GET" summary:"修改密码" tags:"用户"`
	Username    string `v:"required#用户名不能为空"` // 用户名
	NewPassword string `v:"required#新密码不能为空"` // 新密码
}

// ChangePasswordRes /** 修改密码
type ChangePasswordRes struct {
	g.Meta `method:"GET" summary:"修改密码" tags:"用户"`
}
