package requests

type GetGroupsRequest struct {
	Id int `json:"id" binding:"required"`
}

type InsertGroupRequest struct {
	Name      string `json:"name" binding:"required"`
	FacultyId int    `json:"faculty_id" binding:"required"`
}

type UpdateGroupRequest struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteGroupRequest struct {
	Id int `json:"id" binding:"required"`
}
