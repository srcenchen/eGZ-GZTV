// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package upload

import (
	"context"

	"eGZ-GZTV/api/upload/v1"
)

type IUploadV1 interface {
	UploadImage(ctx context.Context, req *v1.UploadImageReq) (res *v1.UploadImageRes, err error)
	UploadVideo(ctx context.Context, req *v1.UploadVideoReq) (res *v1.UploadVideoRes, err error)
}
