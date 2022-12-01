package entity

type JWT struct {
	Entity             `json:"-" db:"-" binding:"-"`
	Token              string `json:"token" binding:"required"`
	ExpirationTimeUnix int64  `json:"expirationTimeUnix" binding:"required"`
	CreateTimeUnix     int64  `json:"createTimeUnix" binding:"required"`
}
