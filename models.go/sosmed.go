package models

type Sosmed struct {
	GormModel
	Nama      string `gorm:"not null" json:"name" form:"name" binding:"required"`
	SosmedURL string `gorm:"nor null" json:"social_media_url" form:"social_media_url" binding:"required"`
	UserID    uint
	User      *User
}
