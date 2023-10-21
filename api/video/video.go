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
	AddVideoGroup(ctx context.Context, req *v1.AddVideoGroupReq) (res *v1.AddVideoGroupRes, err error)
	GetVideoGroupList(ctx context.Context, req *v1.GetVideoGroupListReq) (res *v1.GetVideoGroupListRes, err error)
	GetVideoGroup(ctx context.Context, req *v1.GetVideoGroupReq) (res *v1.GetVideoGroupRes, err error)
	EditVideoDetail(ctx context.Context, req *v1.EditVideoDetailReq) (res *v1.EditVideoDetailRes, err error)
}


