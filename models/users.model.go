package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:password`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
