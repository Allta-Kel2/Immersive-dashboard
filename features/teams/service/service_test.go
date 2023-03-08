package service

import (
	tims "immersiveApp/features/teams"
	"immersiveApp/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)
	expected := tims.TeamEntity{
		Id:        uint(1),
		Name:      "mentor",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	input := tims.TeamEntity{
		Id:        uint(1),
		Name:      "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	srv := New(repo)
	t.Run("Sukses Add", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(expected, nil).Once()

		_, err := srv.Create(expected)
		assert.Nil(t, err)
		assert.Empty(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("eror", func(t *testing.T) {
		_, err := srv.Create(input)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
