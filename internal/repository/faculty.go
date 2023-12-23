package repository

import (
	"context"
	"github.com/IlyaZayats/faculus/internal/entity"
	"github.com/IlyaZayats/faculus/internal/interfaces"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresFacultyRepository struct {
	db *pgxpool.Pool
}

func NewPostgresFacultyRepository(db *pgxpool.Pool) (interfaces.FacultyRepository, error) {
	return &PostgresFacultyRepository{
		db: db,
	}, nil
}

func (r *PostgresFacultyRepository) GetFaculties() ([]entity.Faculty, error) {
	var faculties []entity.Faculty
	q := "SELECT (id, name) FROM Faculties"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return faculties, err
	}
	//faculties, err =
	return r.parseRowsToSlice(rows)

}

func (r *PostgresFacultyRepository) InsertFaculty(faculty entity.Faculty) error {
	q := "INSERT INTO Faculties (name) VALUES ($1)"
	if _, err := r.db.Exec(context.Background(), q, faculty.Name); err != nil {
		return err
	}
	return nil
}

func (r *PostgresFacultyRepository) UpdateFaculty(faculty entity.Faculty) error {
	q := "UPDATE Faculties SET name=$1 WHERE id=$2"
	if _, err := r.db.Exec(context.Background(), q, faculty.Name, faculty.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresFacultyRepository) DeleteFaculty(id int) error {
	q := "DELETE FROM Faculties WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresFacultyRepository) parseRowsToSlice(rows pgx.Rows) ([]entity.Faculty, error) {
	var slice []entity.Faculty
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return slice, err
		}
		slice = append(slice, entity.Faculty{id, name})
	}
	return slice, nil
}
