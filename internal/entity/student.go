package entity

type Student struct {
	Id          int    `db:"id"`
	GroupId     int    `db:"group_id"`
	LastName    string `db:"lastname"`
	FirstName   string `db:"firstname"`
	MiddleName  string `db:"middlename"`
	BirthDate   string `db:"birthdate"`
	PhoneNumber string `db:"phone"`
	Gender      int    `db:"sex"`
}
