package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UploadVideoReq 上传视频文件请求
type UploadVideoReq struct {
	g.Meta `method:"POST" summary:"上传视频文件" tags:"上传"`
	File   *ghttp.UploadFile `json:"file" v:"required#文件不能为空"`
}

// UploadVideoRes 上传视频文件响应
type UploadVideoRes struct {
	FileName string `json:"fileName" example:"121231223.mp4" description:"文件名"`
	Type     string `json:"type" example:"videos" description:"文件类型"`
}
