package entity

type RefreshToken struct {
	Entity             `json:"-" db:"-" binding:"-"`
	Id                 string `json:"id" binding:"required"`
	Token              string `json:"token" binding:"required"`
	ExpirationTimeUnix int64  `json:"expirationTime" binding:"required"`
	CreateTimeUnix     int64  `json:"createTimeUnix" binding:"required"`
}

func NewRefreshToken(id string, token string, expirationTimeUnix int64, createTimeUnix int64) *RefreshToken {
	return &RefreshToken{Id: id, Token: token, ExpirationTimeUnix: expirationTimeUnix, CreateTimeUnix: createTimeUnix}
}
