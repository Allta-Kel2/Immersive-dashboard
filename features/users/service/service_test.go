package service

import (
	"errors"
	"immersiveApp/features/users"
	"immersiveApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewUserDataInterface(t)
	input := users.UserEntity{
		Id:          uint(1),
		TeamId:      1,
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "08123",
		Address:     "pekanbaru",
		Role:        "active",
		Status:      false,
	}
	srv := New(repo)

	mockinput := users.UserEntity{
		Id:          uint(2),
		TeamId:      1,
		FullName:    "",
		Email:       "",
		Password:    "",
		PhoneNumber: "",
		Address:     "",
		Role:        "",
		Status:      false,
	}
	t.Run("succes", func(t *testing.T) {
		repo.On("Store", input).Return(uint(1), nil)
		repo.On("SelectById", uint(1)).Return(input, nil)

		created, err := srv.Create(input)

		assert.NoError(t, err)
		assert.Equal(t, input, created)

		repo.AssertExpectations(t)
	})

	t.Run("eror", func(t *testing.T) {
		_, err := srv.Create(mockinput)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := mocks.NewUserDataInterface(t)
	input := users.UserEntity{
		Id:          uint(1),
		TeamId:      1,
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "08123",
		Address:     "pekanbaru",
		Role:        "active",
		Status:      false,
	}
	srv := New(repo)

	mockList := make([]users.UserEntity, 0)
	mockList = append(mockList, input)

	repo.On("SelectAll").Return(mockList, nil)

	result, err := srv.GetAll()

	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Len(t, result, len(mockList))
	assert.Equal(t, mockList, result)

}
func TestGetById(t *testing.T) {
	repo := mocks.NewUserDataInterface(t)
	input := users.UserEntity{
		Id:          uint(1),
		TeamId:      1,
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "08123",
		Address:     "pekanbaru",
		Role:        "active",
		Status:      false,
	}

	repo.On("SelectById", input.Id).Return(input, nil)

	srv := New(repo)

	result, err := srv.GetById(input.Id)
	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, input, result)
}

func TestDelete(t *testing.T) {
	repo := mocks.NewUserDataInterface(t)
	input := users.UserEntity{
		Id:          uint(1),
		TeamId:      1,
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "08123",
		Address:     "pekanbaru",
		Role:        "active",
		Status:      false,
	}
	t.Run("Sukses Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(input, nil).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(input, errors.New("error")).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NotNil(t, err)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {}

// 	repo := mocks.NewTeamDataInterface(t)
// 	expected := tims.TeamEntity{
// 		Id:        1,
// 		Name:      "team A",
// 		CreatedAt: time.Time{},
// 		UpdatedAt: time.Time{},
// 	}
// 	input := tims.TeamEntity{
// 		Id:        1,
// 		Name:      "team B",
// 		CreatedAt: time.Time{},
// 		UpdatedAt: time.Time{},
// 	}
// 	t.Run("Update success", func(t *testing.T) {
// 		repo.On("Update", expected, 1).Return(input, nil).Once()
// 		srv := New(repo)

// 		res, err := srv.Update(input, 1)
// 		assert.Nil(t, err)
// 		assert.NotEmpty(t, res)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Update Fail", func(t *testing.T) {
// 		repo.On("Update", expected, 1).Return(tims.TeamEntity{}, errors.New("error update")).Once()
// 		srv := New(repo)
// 		res, err := srv.Update(input, 0)
// 		assert.Empty(t, res)
// 		assert.NotNil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// }
