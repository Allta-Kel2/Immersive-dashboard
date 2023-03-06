package delivery

import (
	"immersiveApp/features/teams/delivery"
	"immersiveApp/features/users"
	"reflect"
)

type UserResponse struct {
	Id          uint                   `json:"id"`
	TeamId      uint                   `json:"team_id"`
	FullName    string                 `json:"full_name"`
	Email       string                 `json:"email"`
	Role        string                 `json:"role"`
	Status      bool                   `json:"status"`
	PhoneNumber string                 `json:"phone_number"`
	Address     string                 `json:"address"`
	Team        *delivery.TeamResponse `json:"team,omitempty"`
}

func UserEntityToUserResponse(userEntity users.UserEntity) UserResponse {
	result := UserResponse{
		Id:          userEntity.Id,
		TeamId:      userEntity.TeamId,
		FullName:    userEntity.FullName,
		Email:       userEntity.Email,
		Role:        userEntity.Team.Name,
		Status:      userEntity.Status,
		PhoneNumber: userEntity.PhoneNumber,
		Address:     userEntity.Address,
	}

	if !reflect.ValueOf(userEntity.Team).IsZero() {
		result.Team = &delivery.TeamResponse{
			Id:   userEntity.Team.Id,
			Name: userEntity.Team.Name,
		}
	}

	return result
}

func ListUserEntityToUserResponse(dataCore []users.UserEntity) []UserResponse {
	var dataResponses []UserResponse
	for _, v := range dataCore {
		dataResponses = append(dataResponses, UserEntityToUserResponse(v))
	}
	return dataResponses
}
