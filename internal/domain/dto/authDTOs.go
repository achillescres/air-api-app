package dto

type LoginInput struct {
	Login    string `json:"username" binding:"required, max=20"`
	Password string `json:"password" binding:"required, max=20"`
}

type RegisterInput struct {
	Login    string `json:"username" binding:"required, max=20"`
	Password string `json:"password" binding:"required, max=20"`
}
