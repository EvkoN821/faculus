package interfaces

import "github.com/IlyaZayats/faculus/internal/entity"

type GroupRepository interface {
	GetGroups(id int) ([]entity.Group, error)
	UpdateGroup(group entity.Group) error
	InsertGroup(group entity.Group) error
	DeleteGroup(id int) error
}
