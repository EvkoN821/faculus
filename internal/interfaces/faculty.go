package interfaces

import "github.com/IlyaZayats/faculus/internal/entity"

type FacultyRepository interface {
	GetFaculties() ([]entity.Faculty, error)
	UpdateFaculty(faculty entity.Faculty) error
	InsertFaculty(faculty entity.Faculty) error
	DeleteFaculty(id int) error
}
