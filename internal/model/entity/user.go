package entity

// UserTable /*
type UserTable struct {
	Id       int    `gorm:"primary_key;auto_increment;not null"`
	Username string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}
