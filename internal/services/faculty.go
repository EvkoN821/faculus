package services

import (
	"github.com/IlyaZayats/faculus/internal/entity"
	"github.com/IlyaZayats/faculus/internal/interfaces"
	"strconv"
)

type FacultyService struct {
	repo interfaces.FacultyRepository
}

func NewFacultyService(repo interfaces.FacultyRepository) (*FacultyService, error) {
	return &FacultyService{
		repo: repo,
	}, nil
}

func (s *FacultyService) GetFaculties() ([]map[string]string, error) {
	faculties, err := s.repo.GetFaculties()
	if err != nil {
		return nil, err
	}
	facultiesSlice := []map[string]string{}
	for _, item := range faculties {
		facultiesMap := map[string]string{
			"id":   strconv.Itoa(item.Id),
			"name": item.Name,
		}
		facultiesSlice = append(facultiesSlice, facultiesMap)
	}
	return facultiesSlice, nil
}

func (s *FacultyService) InsertFaculty(name string) error {
	return s.repo.InsertFaculty(entity.Faculty{Id: 0, Name: name})
}

func (s *FacultyService) UpdateFaculty(id int, name string) error {
	return s.repo.UpdateFaculty(entity.Faculty{Id: id, Name: name})
}

func (s *FacultyService) DeleteFaculty(id int) error {
	return s.repo.DeleteFaculty(id)
}
