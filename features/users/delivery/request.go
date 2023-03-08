package delivery

import "immersiveApp/features/users"

type UserRequest struct {
	TeamId      uint   `json:"team_id" form:"team_id" validate:"required"`
	FullName    string `json:"full_name" form:"full_name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
	Role        string `json:"role" form:"role" validate:"required"`
	Status      bool   `json:"status" form:"status"`
}

func UserRequestToUserEntity(userRequest UserRequest) users.UserEntity {
	return users.UserEntity{
		TeamId:      userRequest.TeamId,
		FullName:    userRequest.FullName,
		Email:       userRequest.Email,
		Password:    userRequest.Password,
		PhoneNumber: userRequest.PhoneNumber,
		Address:     userRequest.Address,
		Role:        userRequest.Role,
		Status:      userRequest.Status,
	}
}
