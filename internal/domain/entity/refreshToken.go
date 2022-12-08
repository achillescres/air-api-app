package entity

import "github.com/achillescres/saina-api/pkg/object/oid"

type RefreshToken struct {
	Entity             `json:"-" db:"-" binding:"-"`
	Id                 oid.Id `json:"id" binding:"required"`
	Token              string `json:"token" binding:"required"`
	ExpirationTimeUnix int64  `json:"expirationTime" binding:"required"`
	CreateTimeUnix     int64  `json:"createTimeUnix" binding:"required"`
}

func NewRefreshToken(id oid.Id, token string, expirationTimeUnix int64, createTimeUnix int64) *RefreshToken {
	return &RefreshToken{Id: id, Token: token, ExpirationTimeUnix: expirationTimeUnix, CreateTimeUnix: createTimeUnix}
}
