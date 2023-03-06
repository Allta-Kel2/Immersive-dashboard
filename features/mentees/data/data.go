package data

import (
	"immersiveApp/features/mentees"
	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentees.MenteeDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll() ([]mentees.MenteeEntity, error) {
	var mentees []Mentee
	if err := q.db.Preload("Mentee").Find(&mentees); err.Error != nil {
		return nil, err.Error
	}
	return ListMenteeToMenteeEntity(mentees), nil
}

func (q *query) SelectById(id uint) (mentees.MenteeEntity, error) {
	var mentee Mentee
	if err := q.db.Preload("Mentee").First(&mentee, id); err.Error != nil {
		return mentees.MenteeEntity{}, err.Error
	}
	return MenteeToMenteeEntity(mentee), nil
}

func (q *query) Insert(menteeEntity mentees.MenteeEntity) (uint, error) {
	mentee := MenteeEntitytoClass(menteeEntity)
	if err := q.db.Create(&mentee); err.Error != nil {
		return 0, err.Error
	}
	return mentee.ID, nil
}

func (q *query) Edit(menteeEntity mentees.MenteeEntity, id uint) error {
	mentee := MenteeEntitytoClass(menteeEntity)
	if err := q.db.Where("id", id).Updates(&mentee); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Remove(id uint) error {
	var mentee Mentee
	if err := q.db.Delete(&mentee, id); err.Error != nil {
		return err.Error
	}
	return nil
}