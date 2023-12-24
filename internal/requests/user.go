package requests

type LoginRequest struct {
	User map[string]interface{} `json:"user" binding:"required"`
	//Login string `json:"login" binding:"required"`
	//Pwd   string `json:"pwd" binding:"required"`
}
