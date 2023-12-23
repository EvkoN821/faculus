package interfaces

import "github.com/IlyaZayats/faculus/internal/entity"

type StudentRepository interface {
	GetStudents(id int) ([]entity.Student, error)
	UpdateStudent(student entity.Student) error
	InsertStudent(student entity.Student) error
	DeleteStudent(id int) error
}
