package services

import (
	"github.com/IlyaZayats/faculus/internal/entity"
	"github.com/IlyaZayats/faculus/internal/interfaces"
	"strconv"
)

type StudentService struct {
	repo interfaces.StudentRepository
}

func NewStudentService(repo interfaces.StudentRepository) (*StudentService, error) {
	return &StudentService{
		repo: repo,
	}, nil
}

func (s *StudentService) GetStudents(id int) ([]map[string]string, error) {
	students, err := s.repo.GetStudents(id)
	if err != nil {
		return nil, err
	}
	var studentsSlice []map[string]string
	for _, item := range students {
		studentsMap := map[string]string{
			"id":         strconv.Itoa(item.Id),
			"group_id":   strconv.Itoa(item.GroupId),
			"sex":        strconv.Itoa(item.GroupId),
			"firstname":  item.FirstName,
			"middlename": item.MiddleName,
			"lastname":   item.LastName,
			"phone":      item.PhoneNumber,
			"birthdate":  item.BirthDate,
		}
		studentsSlice = append(studentsSlice, studentsMap)
	}
	return studentsSlice, nil
}

func (s *StudentService) InsertStudent(firstName, lastName, middleName, phone, birthDate string, groupId, sex int) error {
	return s.repo.InsertStudent(entity.Student{Id: 0, FirstName: firstName, LastName: lastName, MiddleName: middleName, PhoneNumber: phone, BirthDate: birthDate, Gender: sex, GroupId: groupId})
}

func (s *StudentService) UpdateStudent(firstName, lastName, middleName, phone, birthDate string, id, sex int) error {
	return s.repo.UpdateStudent(entity.Student{Id: id, FirstName: firstName, LastName: lastName, MiddleName: middleName, PhoneNumber: phone, BirthDate: birthDate, Gender: sex, GroupId: 0})
}

func (s *StudentService) DeleteStudent(id int) error {
	return s.repo.DeleteStudent(id)
}
