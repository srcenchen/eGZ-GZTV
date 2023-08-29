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

// GetPull 获取默认拉流方式
func GetPull() string {
	var settingTable entity.SettingTable
	model.GetDatabase().Where("key = ?", "pull").First(&settingTable)
	return settingTable.Value
}

// SetPull 设置默认拉流方式
func SetPull(pull string) {
	var settingTable entity.SettingTable
	model.GetDatabase().Where("key = ?", "pull").First(&settingTable)
	settingTable.Value = pull
	model.GetDatabase().Save(&settingTable)
}
