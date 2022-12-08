package entity

type JWT struct {
	Entity             `json:"-" db:"-" binding:"-"`
	Token              string `json:"token" binding:"required"`
	ExpirationTimeUnix int64  `json:"expirationTimeUnix" binding:"required"`
	CreateTimeUnix     int64  `json:"createTimeUnix" binding:"required"`
}

func NewJWT(token string, expirationTimeUnix int64, createTimeUnix int64) *JWT {
	return &JWT{Token: token, ExpirationTimeUnix: expirationTimeUnix, CreateTimeUnix: createTimeUnix}
}
