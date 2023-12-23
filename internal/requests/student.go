package requests

type GetStudentsRequest struct {
	Id int `json:"id" binding:"required"`
}

type InsertStudentRequest struct {
	FirstName   string `json:"name" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	MiddleName  string `json:"middlename" binding:"required"`
	BirthDate   string `json:"birthdate" binding:"required"`
	PhoneNumber string `json:"phone" binding:"required"`
	Gender      int    `json:"sex" binding:"required"`
	GroupId     int    `json:"groupid" binding:"required"`
}

type UpdateStudentRequest struct {
	Id          int    `json:"id" binding:"required"`
	FirstName   string `json:"name" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	MiddleName  string `json:"middlename" binding:"required"`
	BirthDate   string `json:"birthdate" binding:"required"`
	PhoneNumber string `json:"phone" binding:"required"`
	Gender      int    `json:"sex" binding:"required"`
}

type DeleteStudentRequest struct {
	Id int `json:"id" binding:"required"`
}
