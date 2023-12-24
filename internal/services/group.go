package services

import (
	"github.com/IlyaZayats/faculus/internal/entity"
	"github.com/IlyaZayats/faculus/internal/interfaces"
	"strconv"
)

type GroupService struct {
	repo interfaces.GroupRepository
}

func NewGroupService(repo interfaces.GroupRepository) (*GroupService, error) {
	return &GroupService{
		repo: repo,
	}, nil
}

func (s *GroupService) GetGroups() ([]map[string]string, error) {
	groups, err := s.repo.GetGroups()
	if err != nil {
		return nil, err
	}
	var groupsSlice []map[string]string
	for _, item := range groups {
		groupsMap := map[string]string{
			"id":         strconv.Itoa(item.Id),
			"faculty_id": strconv.Itoa(item.FacultyId),
			"name":       item.Name,
		}
		groupsSlice = append(groupsSlice, groupsMap)
	}
	return groupsSlice, nil
}

func (s *GroupService) InsertGroup(facultyId int, name string) error {
	return s.repo.InsertGroup(entity.Group{Id: 0, FacultyId: facultyId, Name: name})
}

func (s *GroupService) UpdateGroup(id int, name string) error {
	return s.repo.UpdateGroup(entity.Group{Id: id, FacultyId: 0, Name: name})
}

func (s *GroupService) DeleteGroup(id int) error {
	return s.repo.DeleteGroup(id)
}
