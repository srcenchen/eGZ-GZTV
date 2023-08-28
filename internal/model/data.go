package model

import (
	"eGZ-GZTV/internal/model/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once
var db *gorm.DB

func GetDatabase() *gorm.DB {
	once.Do(func() {
		db, _ = gorm.Open(sqlite.Open("./resource/database/data.db"), &gorm.Config{})
	})
	return db
}

// InitData /**
func InitData() {
	// 自动迁移模式
	if GetDatabase().AutoMigrate(&entity.VideoTable{}, &entity.LiveTable{}, &entity.UserTable{}, &entity.SettingTable{}) != nil {
		panic("failed to migrate database")
	}
	// 判断是否已经初始化过
	var userTable entity.UserTable
	if GetDatabase().Where("id = ?", 1).First(&userTable).RowsAffected == 0 { // 没初始化过
		// 初始化用户
		GetDatabase().Create(&entity.UserTable{Username: "admin", Password: "8dbc6dfc58f9f7cf07eff4bef62c0468"})
		GetDatabase().Create(&entity.UserTable{Username: "sanenchen", Password: "ddaf3b51bac35885fd2d785155daf7b6"})
		// 初始化设置
		GetDatabase().Create(&entity.SettingTable{Key: "notice", Value: "少年不惧岁月长，彼方尚有荣光在。"})
	}
}
