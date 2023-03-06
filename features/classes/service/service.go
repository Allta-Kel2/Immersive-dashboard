package service

import (
	"immersiveApp/features/classes"
)

type ClassService struct {
	Data classes.ClassDataInterface
}

func New(data classes.ClassDataInterface) classes.ClassServiceInterface {
	return &ClassService{
		Data: data,
	}
}

func (s *ClassService) GetAll() ([]classes.ClassEntity, error) {
	return s.Data.SelectAll()
}

func (s *ClassService) GetById(id uint) (classes.ClassEntity, error) {
	return s.Data.SelectById(id)
}

func (s *ClassService) Create(classEntity classes.ClassEntity) (classes.ClassEntity, error) {
	user_id, err := s.Data.Store(classEntity)
	if err != nil {
		return classes.ClassEntity{}, err
	}

	return s.Data.SelectById(user_id)
}

func (s *ClassService) Update(classEntity classes.ClassEntity, id uint) (classes.ClassEntity, error) {
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	err := s.Data.Edit(classEntity, id)
	if err != nil {
		return classes.ClassEntity{}, err
	}
	return s.Data.SelectById(id)
}

func (s *ClassService) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}

	return s.Data.Destroy(id)
}
