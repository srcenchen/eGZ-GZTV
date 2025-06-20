package upload

import (
	"context"
	"eGZ-GZTV/utils"

	"eGZ-GZTV/api/upload/v1"
)

func (c *ControllerV1) UploadImage(_ context.Context, req *v1.UploadImageReq) (res *v1.UploadImageRes, err error) {
	file := req.File
	file.Filename = utils.FileNameEncode(file.Filename)
	fileName, err := file.Save("./resource/upload/images/")

	res = &v1.UploadImageRes{
		Type:     "images",
		FileName: fileName,
	}
	return
}
