package user

import (
	"goFiber/main/config"
	"goFiber/main/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(email string, name string, password string) (models.User, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	var newUser = models.User{Email: email, Name: name, Password: string(hashPassword)}

	config.Database.Db.Create(&newUser)

	return newUser, nil
}
