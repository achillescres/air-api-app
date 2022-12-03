package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/storage"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/db/postgresql"
	"github.com/achillescres/saina-api/pkg/object/oid"
	log "github.com/sirupsen/logrus"
)

type RefreshTokenRepository storage.RefreshTokenStorage

type refreshTokenRepository struct {
	pool postgresql.PGXPool
}

func NewRefreshTokenRepository(pool postgresql.PGXPool) RefreshTokenRepository {
	return &refreshTokenRepository{pool: pool}
}

var _ RefreshTokenRepository = (*refreshTokenRepository)(nil)

func (rTR *refreshTokenRepository) GetById(ctx context.Context, id oid.Id) (*entity.RefreshToken, error) {
	//TODO implement me
	panic("implement me")
}

func (rTR *refreshTokenRepository) GetByToken(ctx context.Context, refreshToken string) (*entity.RefreshToken, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM public.refresh_tokens WHERE token=$1")
	query := rTR.pool.QueryRow(ctx, sqlQuery,
		refreshToken)
	rT := entity.RefreshToken{}
	err := query.Scan(&rT)
	if err != nil {
		return nil, err
	}

	return &rT, nil
}

func (rTR *refreshTokenRepository) GetAll(ctx context.Context) ([]*entity.RefreshToken, error) {
	//TODO implement me
	panic("implement me")
}

func (rTR *refreshTokenRepository) GetAllInMap(ctx context.Context) (map[oid.Id]*entity.RefreshToken, error) {
	//TODO implement me
	panic("implement me")
}

func (rTR *refreshTokenRepository) Store(ctx context.Context, rTC dto.RefreshTokenCreate) (*entity.RefreshToken, error) {
	sqlQuery := fmt.Sprintf("INSERT INTO public.refresh_tokens (token, expiration_time_unix, created_time_unix) VALUES ($1, $2, $3) RETURNING (id)")
	query, err := rTR.pool.Query(ctx, sqlQuery, rTC.Token, rTC.ExpirationTimeUnix, rTC.CreateTimeUnix)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	if !query.Next() {
		err := errors.New("error sql didn't return id of new RefreshToken")
		log.Errorln(err.Error()) // TODO wtf
		return nil, err
	}

	var id string
	err = query.Scan(&id)
	if err != nil {
		log.Errorf("error couldn't scan id from row of new RefreshToken") // TODO wtf
		return nil, err
	}

	return rTC.ToEntity(oid.ToId(id)), nil
}

func (rTR *refreshTokenRepository) DeleteById(ctx context.Context, id oid.Id) (*entity.RefreshToken, error) {
	//TODO implement me
	panic("implement me")
}
