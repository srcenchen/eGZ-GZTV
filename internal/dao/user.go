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
