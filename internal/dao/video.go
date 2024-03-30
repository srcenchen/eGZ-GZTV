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
func AddVideo(title string, description string, videoFile string, headImage string, uploadDate string, groupID int) {
	model.GetDatabase().Create(&entity.VideoTable{Title: title, Description: description, VideoName: videoFile, HeadImage: headImage, UploadDate: uploadDate, GroupId: groupID})
}

// AddGroup 添加分组
func AddGroup(title string, description string, parentGroup string) {
	model.GetDatabase().Create(&entity.GroupTable{Title: title, Description: description, ParentGroup: parentGroup})
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

// DeleteGroupByID 根据id删除视频组
func DeleteGroupByID(id string) {
	var groupTable entity.GroupTable
	model.GetDatabase().Where("id = ?", id).Delete(&groupTable)
}

// GetVideoGroup 获取所有视频组
func GetVideoGroup() []entity.GroupTable {
	var groupTable []entity.GroupTable
	model.GetDatabase().Find(&groupTable)
	return groupTable
}

// GetVideoByGroupID
func GetVideoByGroupID(id string) []entity.VideoTable {
	var videoTable []entity.VideoTable
	model.GetDatabase().Where("group_id = ?", id).Find(&videoTable)
	return videoTable
}

// GetGroupByID 根据id获取视频组
func GetGroupByID(id string) entity.GroupTable {
	var groupTable entity.GroupTable
	model.GetDatabase().Where("id = ?", id).First(&groupTable)
	return groupTable
}

// GetGroupByParent 根据Parent获取视频组
func GetGroupByParent(id string) []entity.GroupTable {
	var groupTable []entity.GroupTable
	model.GetDatabase().Where("parent_group = ?", id).Find(&groupTable)
	return groupTable
}
