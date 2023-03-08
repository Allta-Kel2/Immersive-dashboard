package service

import (
	"immersiveApp/features/users"
	"immersiveApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	data := mocks.NewAuthDataInterface(t)
	srv := New(data)
	mockUser := users.UserEntity{
		Id:       1,
		Email:    "test@example.com",
		Password: "$2a$10$GzLZnYRZuM1J92T8oR0h0eGs/nlPI1RVeX9J0vquSEoEaYq3LRsW2", // "password"
		Role:     "user",
	}

	t.Run("login success", func(t *testing.T) {
		data.On("GetUserByEmailOrId", mockUser.Email, uint(0)).Return(mockUser, nil).Once()
		token, err := srv.Login(mockUser.Email, "password")
		assert.NotNil(t, token)
		assert.Equal(t, err, "GetUserByEmailOrId")
		assert.NotEmpty(t, token)
	})

	t.Run("email and password must be filled", func(t *testing.T) {
		token, err := srv.Login("", "")
		assert.EqualError(t, err, "email and password must be fill")
		assert.Empty(t, token)
	})

	t.Run("user and password not found", func(t *testing.T) {
		token, err := srv.Login("user1@example.com", "")
		assert.EqualError(t, err, "user and password not found")
		assert.Empty(t, token)
	})
}

func TestAuthService_ChangePassword(t *testing.T) {}

// func TestRegister(t *testing.T) {
// 	data := mocks.NewAuthDataInterface(t)
// 	srv := New(data)
// 	user := users.UserEntity{}
// 	t.Run("register success", func(t *testing.T) {
// 		err := srv.Register(users.UserEntity{
// 			Name:     "User 1",
// 			Email:    "user1@example.com",
// 			Password: "password",
// 			Role:     "user",
// 		})
// 		assert.NoError(t, err)
// 		assert.Len(t, user, 1)
// 	})
// }
