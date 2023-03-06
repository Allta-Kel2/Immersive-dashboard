package delivery

import (
	"immersiveApp/features/classes/delivery"
	"immersiveApp/features/mentees"
	"reflect"
)

type MenteeResponse struct {
	Id             uint   `json:"id"`
	ClassId        uint   `json:"class_id"`
	FullName       string `json:"full_name"`
	NickName       string `json:"nick_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	CurrentAddress string `json:"current_address"`
	HomeAddress    string `json:"home_address"`
	Telegram       string `json:"telegram"`
	StatusId       uint   `json:"status_id"`
	// Status          *delivery.StatusResponse
	Gender          string `json:"gender"`
	EducationType   string `json:"education_type"`
	Major           string `json:"major"`
	Graduate        string `json:"graduate"`
	Institution     string `json:"institution"`
	EmergencyName   string `json:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status"`
	Class           *delivery.ClassResponse
}

func MenteeEntitytoMenteeResponse(menteeEntity mentees.MenteeEntity) MenteeResponse {
	result := MenteeResponse{
		Id:              menteeEntity.Id,
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
	if !reflect.ValueOf(menteeEntity.Class).IsZero() {
		result.Class = &delivery.ClassResponse{
			Id: menteeEntity.Class.Id,
			ClassName: menteeEntity.Class.ClassName,
		}
	}
	return result
}

func ListMenteeToMenteeResponse(dataCore []mentees.MenteeEntity) []MenteeResponse{
	var dataResponse []MenteeResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, MenteeEntitytoMenteeResponse(v))
	}
	return dataResponse
}
