package service

import (
	"errors"
	tims "immersiveApp/features/teams"
	"immersiveApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)
	input := tims.TeamEntity{
		Id:   uint(1),
		Name: "",
	}
	srv := New(repo)

	mockTeamEntity := tims.TeamEntity{
		Name: "Team A",
	}
	t.Run("succes", func(t *testing.T) {
		repo.On("Store", mockTeamEntity).Return(uint(1), nil)
		repo.On("SelectById", uint(1)).Return(mockTeamEntity, nil)

		created, err := srv.Create(mockTeamEntity)

		assert.NoError(t, err)
		assert.Equal(t, mockTeamEntity, created)

		repo.AssertExpectations(t)
	})
	t.Run("eror", func(t *testing.T) {
		_, err := srv.Create(input)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)
	mockTeam := tims.TeamEntity{
		Id:   0,
		Name: "team1",
	}

	mockListTeam := make([]tims.TeamEntity, 0)
	mockListTeam = append(mockListTeam, mockTeam)

	repo.On("SelectAll").Return(mockListTeam, nil)

	srv := New(repo)

	result, err := srv.GetAll()

	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Len(t, result, len(mockListTeam))
	assert.Equal(t, mockListTeam, result)

}
func TestGetById(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)
	mockTeam := tims.TeamEntity{
		Id:   1,
		Name: "team1",
	}

	repo.On("SelectById", mockTeam.Id).Return(mockTeam, nil)

	srv := New(repo)

	result, err := srv.GetById(mockTeam.Id)
	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, mockTeam, result)
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)
	input := tims.TeamEntity{
		Id:   1,
		Name: "Team B",
	}

	mockTeamEntity := tims.TeamEntity{
		Name: "Team A",
	}

	t.Run("Sukses Update", func(t *testing.T) {
		repo.On("Update", uint(1), input).Return(mockTeamEntity, nil).Once()
		srv := New(repo)
		res, err := srv.Update(mockTeamEntity, 1)
		assert.NoError(t, err)
		assert.Equal(t, input.Id, res.Id)
		assert.Equal(t, input.Name, res.Name)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Update", func(t *testing.T) {
		repo.On("Update", uint(1), input).Return(tims.TeamEntity{}, errors.New("error update")).Once()
		srv := New(repo)
		var input tims.TeamEntity
		res, err := srv.Update(input, 1)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)

	mockTeamEntity := tims.TeamEntity{
		Name: "Team A",
	}
	t.Run("Sukses Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(mockTeamEntity, nil).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(mockTeamEntity, errors.New("error")).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NotNil(t, err)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})
}
