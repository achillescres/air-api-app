package sto

type LoginUserInput struct {
	Login    string `json:"username" binding:"required, max=20"`
	Password string `json:"password" binding:"required, max=20"`
}

type RegisterUserInput struct {
	Login    string `json:"username" binding:"required, max=20"`
	Password string `json:"password" binding:"required, max=20"`
}
