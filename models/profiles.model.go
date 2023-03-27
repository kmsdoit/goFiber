package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Relation        string `json:"relation"`
	Name            string `json:"name"`
	Birth           string `json:"birth"`
	Gender          string `json:"gender"`
	Education       string `json:"education"`
	EducationStatus string `json:"educationStatus"`
	Country         string `json:"country"`
	City            string `json:"city"`
	Occupation      string `json:"occupation"`
	UserId          int    `json:"userId"`
}
