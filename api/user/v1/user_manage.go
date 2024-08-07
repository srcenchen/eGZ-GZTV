package v1

import (
	"eGZ-GZTV/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// UserListReq /** 用户列表
type UserListReq struct {
	g.Meta `method:"GET" summary:"用户列表" tags:"用户"`
}

// UserListRes /** 用户列表
type UserListRes struct {
	g.Meta `method:"GET" summary:"用户列表" tags:"用户"`
	List   []entity.UserTable `json:"list"` // 用户列表
}

// AddUserReq /** 添加用户
type AddUserReq struct {
	g.Meta   `method:"POST" summary:"添加用户" tags:"用户"`
	Username string `v:"required#Username不能为空"`
	Password string `v:"required#Username不能为空"`
}

// AddUserRes /** 添加用户
type AddUserRes struct {
	g.Meta `method:"POST" summary:"添加用户" tags:"用户"`
}

// DelUserReq /** 删除用户
type DelUserReq struct {
	g.Meta `method:"POST" summary:"删除用户" tags:"用户"`
	Id     string `v:"required#id不为空"`
}

// DelUserRes /** 删除用户
type DelUserRes struct {
	g.Meta `method:"POST" summary:"删除用户" tags:"用户"`
}
