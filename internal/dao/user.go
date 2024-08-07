package dao

import (
	"eGZ-GZTV/internal/model"
	"eGZ-GZTV/internal/model/entity"
)

func VerifyUser(username string, password string) bool {
	var userTable entity.UserTable
	result := model.GetDatabase().Where("username = ?", username).Where("password = ?", password).First(&userTable)
	return result.RowsAffected != 0
}

func ChangePassword(username string, password string) {
	var userTable entity.UserTable
	model.GetDatabase().Where("username = ?", username).First(&userTable)
	userTable.Password = password
	model.GetDatabase().Save(&userTable)
}

// GetUserList 获取所有用户
func GetUserList() []entity.UserTable {
	var groupTable []entity.UserTable
	model.GetDatabase().Find(&groupTable)
	return groupTable
}

// AddUser 添加分组
func AddUser(username string, password string) {
	model.GetDatabase().Create(&entity.UserTable{Username: username, Password: password})
}

// DeleteUserByID 根据id删除User
func DeleteUserByID(id string) {
	var userTable entity.UserTable
	model.GetDatabase().Where("id = ?", id).Delete(&userTable)
}
