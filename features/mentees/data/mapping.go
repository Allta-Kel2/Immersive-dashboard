package data

import (
	"immersiveApp/features/classes"
	"immersiveApp/features/mentees"
	"reflect"
)

func MenteeEntitytoClass(menteeEntity mentees.MenteeEntity) Mentee {
	return Mentee{
		ClassId:         menteeEntity.ClassId,
		FullName:        menteeEntity.FullName,
		NickName:        menteeEntity.NickName,
		Email:           menteeEntity.Email,
		Phone:           menteeEntity.Phone,
		CurrentAddress:  menteeEntity.CurrentAddress,
		HomeAddress:     menteeEntity.HomeAddress,
		Telegram:        menteeEntity.Telegram,
		StatusId:        menteeEntity.StatusId,
		Gender:          menteeEntity.Gender,
		EducationType:   menteeEntity.EducationType,
		Major:           menteeEntity.Major,
		Graduate:        menteeEntity.Graduate,
		Institution:     menteeEntity.Institution,
		EmergencyName:   menteeEntity.EmergencyName,
		EmergencyPhone:  menteeEntity.EmergencyPhone,
		EmergencyStatus: menteeEntity.EmergencyStatus,
	}
}

func MenteeToMenteeEntity(mentee Mentee) mentees.MenteeEntity {
	result := mentees.MenteeEntity{
		Id:             mentee.ID,
		ClassId:        mentee.ClassId,
		FullName:       mentee.FullName,
		NickName:       mentee.NickName,
		Email:          mentee.Email,
		Phone:          mentee.Phone,
		CurrentAddress: mentee.CurrentAddress,
		HomeAddress:    mentee.HomeAddress,
		Telegram:       mentee.Telegram,
		StatusId:       mentee.StatusId,
		// Status          statuses.StatusEntity
		Gender:          mentee.Gender,
		EducationType:   mentee.EducationType,
		Major:           mentee.Major,
		Graduate:        mentee.Graduate,
		Institution:     mentee.Institution,
		EmergencyName:   mentee.EmergencyName,
		EmergencyPhone:  mentee.EmergencyPhone,
		EmergencyStatus: mentee.EmergencyStatus,
		CreatedAt:       mentee.CreatedAt,
		UpdatedAt:       mentee.UpdatedAt,
	}
	if !reflect.ValueOf(mentee.Class).IsZero() {
		result.Class = classes.ClassEntity{
			Id:        mentee.Class.ID,
			ClassName: mentee.Class.ClassName,
		}
	}
	return result
}

func ListMenteeToMenteeEntity(mentee []Mentee) []mentees.MenteeEntity {
	var menteeEntity []mentees.MenteeEntity
	for _, v := range mentee {
		menteeEntity = append(menteeEntity, MenteeToMenteeEntity(v))
	}
	return menteeEntity
}
