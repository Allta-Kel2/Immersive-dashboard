package service

import (
	login "immersiveApp/features/auth/delivery"
	"immersiveApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	data := mocks.NewAuthDataInterface(t)
	srv := New(data)
	isi := login.LoginRequest{Email: "", Password: ""}
	input := login.LoginRequest{Email: "rischi@gmail.com", Password: "123$"}
	t.Run("Success login", func(t *testing.T) {
		data.On("Login", mock.Anything).Return(input, nil).Once()

		_, err := srv.Login(input.Email, input.Password)
		assert.Nil(t, err)
		assert.Empty(t, err)
		data.AssertExpectations(t)
	})
	t.Run("email and password must be fill", func(t *testing.T) {
		_, err := srv.Login(isi.Email, isi.Password)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		data.AssertExpectations(t)
	})
}
