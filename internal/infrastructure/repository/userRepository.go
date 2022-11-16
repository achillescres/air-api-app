package repository

import (
	"api-app/internal/config"
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/internal/domain/storage/dto"
	"api-app/pkg/db/postgresql"
	"api-app/pkg/object/oid"
	"api-app/pkg/security/passlib"
	"context"
)

type UserRepository interface {
	storage.Storage[entity.User, entity.UserView, dto.UserCreate]
}

type userRepository struct {
	pool postgresql.Pool
	cfg  config.DBConfig
}

var _ UserRepository = (*userRepository)(nil)

func NewUserRepository(pool postgresql.Pool, cfg config.DBConfig) UserRepository {
	return &userRepository{pool: pool, cfg: cfg}
}

func (uRepo *userRepository) GetById(ctx context.Context, id oid.Id) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uRepo *userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uRepo *userRepository) Store(ctx context.Context, uC dto.UserCreate) (entity.User, error) {
	hashedPassword, err := passlib.HashPassword(uC.Password)
	if err != nil {
		return entity.User{}, err
	}
	query, err := uRepo.pool.Query(
		ctx,
		"INSERT INTO public.users (login, hashed_password) VALUES ($1, $2) RETURNING (id)",
		uC.Login,
		hashedPassword,
	)
	if err != nil {
		return entity.User{}, err
	}

	if !query.Next() {
		return // TODO WHAT TO DO WTF???!!!?
	}
	var id string
	err := query.Scan(&id)
	if err != nil {
		return // TODO WHAT TO DO WTF??!?!?!?!?
	}

	return entity.User{
		Id: oid.ToId(id),
		View: entity.UserView{
			Login:          uC.Login,
			HashedPassword: hashedPassword,
		},
	}, err
}

func (uRepo *userRepository) DeleteById(ctx context.Context, id oid.Id) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}
