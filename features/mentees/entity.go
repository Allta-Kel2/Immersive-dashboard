package mentees

import (
	"immersiveApp/features/classes"
	"immersiveApp/features/statuses"
	"time"
)

type MenteeEntity struct {
	Id             uint
	ClassId        uint
	Class          classes.ClassEntity
	FullName       string
	NickName       string
	Email          string
	Phone          string
	CurrentAddress string
	HomeAddress    string
	Telegram       string
	StatusId       uint
	Status          statuses.StatusEntity
	Gender          string
	EducationType   string
	Major           string
	Graduate        string
	Institution     string
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MenteeServiceInterface interface {
	GetAll() ([]MenteeEntity, error)
	GetById(id uint) (MenteeEntity, error)
	Create(menteeEntity MenteeEntity) (MenteeEntity, error)
	Update(menteeEntity MenteeEntity, id uint) (MenteeEntity, error)
	Delete(id uint) error
}

type MenteeDataInterface interface {
	SelectAll() ([]MenteeEntity, error)
	SelectById(id uint) (MenteeEntity, error)
	Insert(menteeEntity MenteeEntity) (uint, error)
	Edit(menteeEntity MenteeEntity, id uint) error
	Remove(id uint) error
}
