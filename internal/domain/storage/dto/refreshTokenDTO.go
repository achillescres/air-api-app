package dto

import (
	"api-app/internal/domain/entity"
	"api-app/pkg/object/oid"
)

type RefreshTokenCreate struct {
	Create             `json:"-" db:"-" binding:"-"`
	Token              string `json:"token" binding:"required"`
	ExpirationTimeUnix int64  `json:"expirationTime" binding:"required"`
	CreateTimeUnix     int64  `json:"createTimeUnix" binding:"required"`
}

func NewRefreshTokenCreate(token string, createTimeUnix int64, expirationTimeUnix int64) *RefreshTokenCreate {
	return &RefreshTokenCreate{Token: token, CreateTimeUnix: createTimeUnix, ExpirationTimeUnix: expirationTimeUnix}
}

func (rTC RefreshTokenCreate) ToEntity(id oid.Id) *entity.RefreshToken {
	return &entity.RefreshToken{
		Token:              rTC.Token,
		ExpirationTimeUnix: rTC.ExpirationTimeUnix,
		CreateTimeUnix:     rTC.CreateTimeUnix,
	}
}
