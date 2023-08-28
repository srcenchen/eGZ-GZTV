// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package live

import (
	"context"
	
	"eGZ-GZTV/api/live/v1"
)

type ILiveV1 interface {
	GetLiveList(ctx context.Context, req *v1.GetLiveListReq) (res *v1.GetLiveListRes, err error)
	GetLive(ctx context.Context, req *v1.GetLiveReq) (res *v1.GetLiveRes, err error)
	AddLive(ctx context.Context, req *v1.AddLiveReq) (res *v1.AddLiveRes, err error)
	DeleteLive(ctx context.Context, req *v1.DeleteLiveReq) (res *v1.DeleteLiveRes, err error)
}


