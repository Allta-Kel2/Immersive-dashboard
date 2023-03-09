package service

// import (
// 	"immersiveApp/features/mentees"
// 	"immersiveApp/mocks"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCreate(t *testing.T) {
// 	// Create a mock object for MenteeServiceInterface
// 	mockMenteeService := new(mocks.MenteeServiceInterface)
// 	// Set up an expectation for the Create method
// 	mockMentee := mentees.MenteeEntity{
// 		Id:              uint(1),
// 		ClassId:         1,
// 		FullName:        "meong milanium",
// 		NickName:        "meong",
// 		Email:           "miyong@mail",
// 		Phone:           "089765",
// 		CurrentAddress:  "jl. sesama",
// 		HomeAddress:     "malang",
// 		Telegram:        "youma",
// 		StatusId:        1,
// 		Gender:          "L",
// 		EducationType:   "non-informatics",
// 		Major:           "DKV",
// 		Graduate:        "1040",
// 		Institution:     "UniGa",
// 		EmergencyName:   "Jola",
// 		EmergencyPhone:  "089890",
// 		EmergencyStatus: "Kakek",
// 	}
// 	mockMenteeService.On("Create", mockMentee).Return(mockMentee, nil)

// 	srv := New(mockMenteeService)
// 	// Call the function that uses the mock object
// 	result, err := srv.Create(mockMentee)

// 	// Assert that the expectations were met
// 	mockMenteeService.AssertExpectations(t)
// 	assert.Nil(t, err)
// 	assert.Equal(t, mockMentee, result)
// }

// // func TestGetAll(t *testing.T) {

// // }

// // func TestGetById(t *testing.T) {
// // 	repo := mocks.NewMenteeDataInterface(t)
// // 	input := mentees.MenteeEntity{
// // 		Id:              uint(1),
// // 		ClassId:         1,
// // 		FullName:        "meong milanium",
// // 		NickName:        "meong",
// // 		Email:           "miyong@mail",
// // 		Phone:           "089765",
// // 		CurrentAddress:  "jl. sesama",
// // 		HomeAddress:     "malang",
// // 		Telegram:        "youma",
// // 		StatusId:        1,
// // 		Gender:          "L",
// // 		EducationType:   "non-informatics",
// // 		Major:           "DKV",
// // 		Graduate:        "1040",
// // 		Institution:     "UniGa",
// // 		EmergencyName:   "Jola",
// // 		EmergencyPhone:  "089890",
// // 		EmergencyStatus: "Kakek",
// // 	}
// // 	repo.On("SelectById", input.Id).Return(input, nil)

// // 	srv := New(repo)

// // 	result, err := srv.GetById(input.Id)
// // 	repo.AssertExpectations(t)
// // 	assert.NoError(t, err)
// // 	assert.Equal(t, input, result)
// // }

// // func TestGetFeedbackById(t *testing.T) {

// // }

// // func TestUpdate(t *testing.T) {

// // }

// // func TestDelete(t *testing.T) {
// // 	repo := mocks.NewMenteeDataInterface(t)
// // 	input := mentees.MenteeEntity{
// // 		Id:              uint(1),
// // 		ClassId:         1,
// // 		FullName:        "meong milanium",
// // 		NickName:        "meong",
// // 		Email:           "miyong@mail",
// // 		Phone:           "089765",
// // 		CurrentAddress:  "jl. sesama",
// // 		HomeAddress:     "malang",
// // 		Telegram:        "youma",
// // 		StatusId:        1,
// // 		Gender:          "L",
// // 		EducationType:   "non-informatics",
// // 		Major:           "DKV",
// // 		Graduate:        "1040",
// // 		Institution:     "UniGa",
// // 		EmergencyName:   "Jola",
// // 		EmergencyPhone:  "089890",
// // 		EmergencyStatus: "Kakek",
// // 	}
// // 	t.Run("Sukses Delete", func(t *testing.T) {
// // 		repo.On("Delete", mock.Anything).Return(input, nil).Once()
// // 		srv := New(repo)
// // 		err := srv.Delete(1)
// // 		assert.Nil(t, err)
// // 		assert.NotEmpty(t, err)
// // 		repo.AssertExpectations(t)
// // 	})
// // 	t.Run("Gagal Delete", func(t *testing.T) {
// // 		repo.On("Delete", mock.Anything).Return(input, errors.New("error")).Once()
// // 		srv := New(repo)
// // 		err := srv.Delete(1)
// // 		assert.NotNil(t, err)
// // 		assert.NotEmpty(t, err)
// // 		repo.AssertExpectations(t)
// // 	})
// // }
