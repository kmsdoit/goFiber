package user

import (
	"goFiber/main/config"
	"goFiber/main/models"
	"goFiber/main/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateUserService(email string, name string, password string) (models.User, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	var newUser = models.User{Email: email, Name: name, Password: string(hashPassword)}

	config.Database.Db.Create(&newUser)

	return newUser, nil
}

func LoginUserService(email string, password string) (string, error) {
	loginInfo := models.User{
		Email:    email,
		Password: password,
	}

	userByEmail, isUserExist, _ := IsUserExistByEmail(email)

	if isUserExist == false {
		return "", nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(password))

	if err != nil {
		return "password error", nil
	}

	t, _, err := utils.CreateJWTToken(loginInfo)

	if err != nil {
		return "token generate error", err
	}

	return t, nil
}

func IsUserExistByEmail(email string) (models.User, bool, error) {

	loginInfo := models.User{
		Email: email,
	}

	result := config.Database.Db.Find(&loginInfo, "email = ?", email)

	if result.RowsAffected == 0 {
		return models.User{}, false, nil
	}

	log.Println(&loginInfo)

	return loginInfo, true, nil

}

func GetAllUserService() ([]models.User, error) {
	var Users []models.User

	config.Database.Db.Find(&Users)

	return Users, nil
}
