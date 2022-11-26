package entity

type JWT struct {
	Token string `json:"token"`
}

type RefreshToken struct {
	Token string `json:'token'`
}
