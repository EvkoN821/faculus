package entity

type Group struct {
	Id        int    `db:"id"`
	FacultyId int    `db:"faculty_id"`
	Name      string `db:"name"`
}
