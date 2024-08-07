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

	// 是否存在标题
	result = model.GetDatabase().Where("key = ?", "title").First(&settingTable)
	if result.RowsAffected == 0 {
		// 初始化设置
		model.GetDatabase().Create(&entity.SettingTable{Key: "title", Value: "赣中电视台"})
	}
	// 是否存在页脚
	result = model.GetDatabase().Where("key = ?", "footer").First(&settingTable)
	if result.RowsAffected == 0 {
		// 初始化设置
		model.GetDatabase().Create(&entity.SettingTable{Key: "footer", Value: "主办方：江苏省赣榆高级中学经济开发区校区"})
	}
}
