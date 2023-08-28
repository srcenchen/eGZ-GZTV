// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package setting

import (
	"context"
	
	"eGZ-GZTV/api/setting/v1"
)

type ISettingV1 interface {
	GetNotice(ctx context.Context, req *v1.GetNoticeReq) (res *v1.GetNoticeRes, err error)
	SetNotice(ctx context.Context, req *v1.SetNoticeReq) (res *v1.SetNoticeRes, err error)
}


