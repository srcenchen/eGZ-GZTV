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

// GetTitle 获取默认
func GetTitle() string {
	var settingTable entity.SettingTable
	model.GetDatabase().Where("key = ?", "title").First(&settingTable)
	return settingTable.Value
}

// SetTitle 设置
func SetTitle(title string) {
	var settingTable entity.SettingTable
	model.GetDatabase().Where("key = ?", "title").First(&settingTable)
	settingTable.Value = title
	model.GetDatabase().Save(&settingTable)
}

// GetFooter 获取默认
func GetFooter() string {
	var settingTable entity.SettingTable
	model.GetDatabase().Where("key = ?", "footer").First(&settingTable)
	return settingTable.Value
}

// SetFooter 设置
func SetFooter(footer string) {
	var settingTable entity.SettingTable
	model.GetDatabase().Where("key = ?", "footer").First(&settingTable)
	settingTable.Value = footer
	model.GetDatabase().Save(&settingTable)
}
