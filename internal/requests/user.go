package requests

type LoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"`
	UserType string `json:"name" binding:"required"`
}
