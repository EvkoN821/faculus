package requests

type InsertFacultyRequest struct {
	Faculty map[string]interface{} `json:"faculty" binding:"required"`
	//Name string `json:"name" binding:"required"`
}

type UpdateFacultyRequest struct {
	Faculty map[string]interface{} `json:"faculty" binding:"required"`
	//Id   int    `json:"id" binding:"required"`
	//Name string `json:"name" binding:"required"`
}

type DeleteFacultyRequest struct {
	Id int `json:"id" binding:"required"`
}
