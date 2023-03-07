package service

import "immersiveApp/features/mentees"

type menteeService struct {
	Data mentees.MenteeDataInterface
}

func New(data mentees.MenteeDataInterface) mentees.MenteeServiceInterface {
	return &menteeService{
		Data: data,
	}
}

func (s *menteeService) GetAll() ([]mentees.MenteeEntity, error) {
	return s.Data.SelectAll()
}

func (s *menteeService) GetById(id uint) (mentees.MenteeEntity, error) {
	return s.Data.SelectById(id)
}

func (s *menteeService) Create(mentee mentees.MenteeEntity) (mentees.MenteeEntity, error) {
	user_id, err := s.Data.Insert(mentee)
	if err != nil {
		return mentees.MenteeEntity{}, err
	}

	return s.Data.SelectById(user_id)
}

func (s *menteeService) Update(mentee mentees.MenteeEntity, id uint) (mentees.MenteeEntity, error) {
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	err := s.Data.Edit(mentee, id)
	if err != nil {
		return mentees.MenteeEntity{}, err
	}

	return s.Data.SelectById(id)
}

func (s *menteeService) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}

	return s.Data.Remove(id)
}
