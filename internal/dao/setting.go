package dao

import (
	"eGZ-GZTV/internal/model"
	"eGZ-GZTV/internal/model/entity"
)

// GetNotice 获取公告
func GetNotice() string {
	var settingTable entity.SettingTable
	model.GetDatabase().Where("key = ?", "notice").First(&settingTable)
	return settingTable.Value
}

// SetNotice 设置公告
func SetNotice(notice string) {
	var settingTable entity.SettingTable
	model.GetDatabase().Where("key = ?", "notice").First(&settingTable)
	settingTable.Value = notice
	model.GetDatabase().Save(&settingTable)
}
