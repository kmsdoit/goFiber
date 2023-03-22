package models

import "gorm.io/gorm"

type Bookmark struct {
	gorm.Model
	Name      string `json:"name"`
	Url       string `json:"url"`
	UserRefer int    `json:"user_id"`
	User      User   `gorm:"foreignKey:UserRefer"`
}
