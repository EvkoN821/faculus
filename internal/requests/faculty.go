package requests

type InsertFacultyRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateFacultyRequest struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteFacultyRequest struct {
	Id int `json:"id" binding:"required"`
}
