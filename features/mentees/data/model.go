package data

import (
	c "immersiveApp/features/classes/data"
	s "immersiveApp/features/statuses/data"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	ClassId         int
	Class           *c.Class `gorm:"foreignKey:ClassId"`
	FullName        string
	NickName        string
	Email           string
	Phone           string
	CurrentAddress  string
	HomeAddress     string
	Telegram        string
	StatusId        int
	Status          *s.Status `gorm:"foreignKey:StatusId"`
	Gender          string
	EducationType   string
	Major           string
	Graduate        string
	Institution     string
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string
}
