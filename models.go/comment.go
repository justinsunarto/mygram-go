package models

type Comment struct {
	GormModel
	PhotoID int    `gorm:"not null" json:"photo_id" form:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message"`
	UserID  uint
	User    *User
}
