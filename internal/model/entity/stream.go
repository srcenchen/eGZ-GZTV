package entity

// VideoTable /*
type VideoTable struct {
	Id          int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Title       string `gorm:"column:title;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:varchar(255);not null"`
	VideoName   string `gorm:"column:video_name;type:varchar(255);not null"`
	HeadImage   string `gorm:"column:head_image;type:varchar(255);not null"`
	UploadDate  string `gorm:"column:upload_date;type:dateTime;not null"`
	ViewCount   int    `gorm:"column:view_count;type:int;not null"`
}

// LiveTable /*
type LiveTable struct {
	Id          int    `gorm:"primary_key;auto_increment;not null"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	AppName     string `gorm:"type:varchar(255)"`
	StreamName  string `gorm:"type:varchar(255);not null"`
	HeadImage   string `gorm:"type:varchar(255);not null"`
	SubmitDate  string `gorm:"type:dateTime;not null"`
}
