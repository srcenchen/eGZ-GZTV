package dao

import (
	"eGZ-GZTV/internal/model"
	"eGZ-GZTV/internal/model/entity"
)

// GetLives 获取所有直播列表
func GetLives() []entity.LiveTable {
	var liveTable []entity.LiveTable
	model.GetDatabase().Find(&liveTable)
	return liveTable
}

// AddLive 添加直播
func AddLive(title string, appName string, streamName string, headImage string, description string, submitDate string) {
	model.GetDatabase().Create(
		&entity.LiveTable{Title: title, Description: description, AppName: appName, StreamName: streamName,
			HeadImage: headImage, SubmitDate: submitDate, LiveState: false})
}

// GetLiveByID 根据id获取直播
func GetLiveByID(id string) entity.LiveTable {
	var liveTable entity.LiveTable
	model.GetDatabase().Where("id = ?", id).First(&liveTable)
	return liveTable
}

// DeleteLiveByID 根据id删除直播
func DeleteLiveByID(id string) {
	var liveTable entity.LiveTable
	model.GetDatabase().Where("id = ?", id).Delete(&liveTable)
}
