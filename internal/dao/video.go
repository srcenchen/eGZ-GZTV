package dao

import (
	"eGZ-GZTV/internal/model"
	"eGZ-GZTV/internal/model/entity"
)

// GetVideos 获取所有视频列表
func GetVideos() []entity.VideoTable {
	var videoTable []entity.VideoTable
	model.GetDatabase().Find(&videoTable)
	return videoTable
}

// AddVideo 添加视频
func AddVideo(title string, description string, videoFile string, headImage string, uploadDate string) {
	model.GetDatabase().Create(&entity.VideoTable{Title: title, Description: description, VideoName: videoFile, HeadImage: headImage, UploadDate: uploadDate})
}

// GetVideoByID 根据id获取视频
func GetVideoByID(id string) entity.VideoTable {
	var videoTable entity.VideoTable
	model.GetDatabase().Where("id = ?", id).First(&videoTable)
	return videoTable
}

// DeleteVideoByID 根据id删除视频
func DeleteVideoByID(id string) {
	var videoTable entity.VideoTable
	model.GetDatabase().Where("id = ?", id).Delete(&videoTable)
}
