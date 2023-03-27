package profile

import (
	"goFiber/main/config"
	"goFiber/main/models"
)

func CreateProfileService(newProfile models.Profile) (models.Profile, error) {
	config.Database.Db.Create(&newProfile)

	return newProfile, nil
}
