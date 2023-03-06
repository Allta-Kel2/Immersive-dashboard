package data

import (
	m "immersiveApp/features/mentees/data"
	s "immersiveApp/features/statuses/data"
	u "immersiveApp/features/users/data"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Notes    string
	Proof    string
	UserId   int
	User     *u.User `gorm:"foreignKey:UserId"`
	MenteeId int
	Mentee   *m.Mentee `gorm:"foreignKey:UserId"`
	StatusId int
	Status   *s.Status `gorm:"foreignKey:UserId"`
	Approved bool
}
