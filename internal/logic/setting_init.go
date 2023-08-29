package logic

import (
	"eGZ-GZTV/internal/model"
	"eGZ-GZTV/internal/model/entity"
)

func SettingInit() {
	var settingTable entity.SettingTable
	// 是否存在公告
	result := model.GetDatabase().Where("key = ?", "notice").First(&settingTable)
	if result.RowsAffected == 0 {
		// 初始化设置
		model.GetDatabase().Create(&entity.SettingTable{Key: "notice", Value: "少年不惧岁月长，彼方尚有荣光在。"})
	}
	// 是否存在默认拉流方式
	result = model.GetDatabase().Where("key = ?", "pull").First(&settingTable)
	if result.RowsAffected == 0 {
		// 初始化设置
		model.GetDatabase().Create(&entity.SettingTable{Key: "pull", Value: "HLS-M3U8"})
	}
}
