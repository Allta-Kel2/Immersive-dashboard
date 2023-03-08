package service

import (
	tims "immersiveApp/features/teams"
	"immersiveApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
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
	repo.On("Store", mockTeamEntity).Return(uint(1), nil)
	repo.On("SelectById", uint(1)).Return(mockTeamEntity, nil)

	created, err := srv.Create(mockTeamEntity)

	assert.Nil(t, err)
	assert.Equal(t, mockTeamEntity, created)

	repo.AssertExpectations(t)

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
	mockTeamEntity := tims.TeamEntity{
		Name: "Team A",
	}

	repo.On("Update", mockTeamEntity.Id).Return(mockTeamEntity, nil).Once()

	srv := New(repo)

	result, err := srv.Update(tims.TeamEntity{Name: "New Team Name"}, uint(1))

	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.Id)
	assert.Equal(t, "Update", result.Name)
}

// func TestDelete(t *testing.T) {
// 	repo := mocks.NewTeamDataInterface(t)
// 	srv := New(repo)
// 	t.Run("Delete Success", func(t *testing.T) {
// 		repo.On("Delete", 1, 1).Return(nil).Once()

// 		_, token := helpers.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		err := srv.Delete(pToken, 1, 1)

// 		assert.Nil(t, err)

// 		repo.AssertExpectations(t)
// 	})
// }
