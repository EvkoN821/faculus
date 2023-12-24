package requests

//type GetStudentsRequest struct {
//	Id int `json:"id" binding:"required"`
//}

type InsertStudentRequest struct {
	Student map[string]interface{} `json:"student" binding:"required"`
	//FirstName   string `json:"firstname" binding:"required"`
	//LastName    string `json:"lastname" binding:"required"`
	//MiddleName  string `json:"middlename" binding:"required"`
	//BirthDate   string `json:"birthdate" binding:"required"`
	//PhoneNumber string `json:"phone" binding:"required"`
	//Gender      int    `json:"sex" binding:"required"`
	//GroupId     int    `json:"group_id" binding:"required"`
}

type UpdateStudentRequest struct {
	Student map[string]interface{} `json:"student" binding:"required"`
	//Id          int    `json:"id" binding:"required"`
	//FirstName   string `json:"firstname" binding:"required"`
	//LastName    string `json:"lastname" binding:"required"`
	//MiddleName  string `json:"middlename" binding:"required"`
	//BirthDate   string `json:"birthdate" binding:"required"`
	//PhoneNumber string `json:"phone" binding:"required"`
	//Gender      int    `json:"sex" binding:"required"`
}

type DeleteStudentRequest struct {
	Id int `json:"id" binding:"required"`
}
