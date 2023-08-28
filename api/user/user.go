// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package user

import (
	"context"
	
	"eGZ-GZTV/api/user/v1"
)

type IUserV1 interface {
	LoginVerify(ctx context.Context, req *v1.LoginVerifyReq) (res *v1.LoginVerifyRes, err error)
	ChangePassword(ctx context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error)
}


