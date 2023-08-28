// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package video

import (
	"context"
	
	"eGZ-GZTV/api/video/v1"
)

type IVideoV1 interface {
	GetVideoList(ctx context.Context, req *v1.GetVideoListReq) (res *v1.GetVideoListRes, err error)
	GetVideo(ctx context.Context, req *v1.GetVideoReq) (res *v1.GetVideoRes, err error)
	AddVideo(ctx context.Context, req *v1.AddVideoReq) (res *v1.AddVideoRes, err error)
	DeleteVideo(ctx context.Context, req *v1.DeleteVideoReq) (res *v1.DeleteVideoRes, err error)
}


