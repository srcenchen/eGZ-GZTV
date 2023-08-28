package upload

import (
	"context"
	"eGZ-GZTV/utils"

	"eGZ-GZTV/api/upload/v1"
)

func (c *ControllerV1) UploadVideo(_ context.Context, req *v1.UploadVideoReq) (res *v1.UploadVideoRes, err error) {
	file := req.File
	file.Filename = utils.FileNameEncode(file.Filename)
	fileName, err := file.Save("./resource/upload/videos/")
	res = &v1.UploadVideoRes{
		Type:     "videos",
		FileName: fileName,
	}
	return
}
