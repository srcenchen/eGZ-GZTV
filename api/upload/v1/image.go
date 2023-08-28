package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UploadImageReq 上传图片文件请求
type UploadImageReq struct {
	g.Meta `method:"POST" summary:"上传图片文件" tags:"上传"`
	File   *ghttp.UploadFile `json:"file" v:"required#文件不能为空"`
}

// UploadImageRes 上传图片文件响应
type UploadImageRes struct {
	FileName string `json:"fileName" example:"121231223.png" description:"文件名"`
	Type     string `json:"type" example:"images" description:"文件类型"`
}
