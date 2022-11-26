package repository

import (
	"api-app/internal/config"
	"api-app/internal/domain/dto"
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/pkg/db/postgresql"
	"api-app/pkg/object/oid"
	"api-app/pkg/security/passlib"
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
)

type UserRepository interface {
	storage.Storage[entity.User, entity.UserView, dto.UserCreate]
	GetByLoginAndHashedPassword(ctx context.Context, login, hashedPassword string) (*entity.User, error)
}

type userRepository struct {
	pool postgresql.PGXPool
	cfg  config.PostgresConfig
}

var _ UserRepository = (*userRepository)(nil)

func NewUserRepository(pool postgresql.PGXPool, cfg config.PostgresConfig) UserRepository {
	return &userRepository{pool: pool, cfg: cfg}
}

func (uRepo *userRepository) GetByLoginAndHashedPassword(
	ctx context.Context,
	login, hashedPassword string,
) (*entity.User, error) {
	sql := "SELECT * FROM public.users WHERE login=$1 AND hashed_password=$2"
	row := uRepo.pool.QueryRow(ctx, sql, login, hashedPassword)
	user := &entity.User{}
	err := row.Scan(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uRepo *userRepository) GetById(ctx context.Context, id oid.Id) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uRepo *userRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uRepo *userRepository) Store(ctx context.Context, uC dto.UserCreate) (*entity.User, error) {
	hashedPassword, err := passlib.Hash(uC.Password)
	if err != nil {
		return &entity.User{}, err
	}
	rows, err := uRepo.pool.Query(
		ctx,
		"INSERT INTO public.users (login, hashed_password) VALUES ($1, $2) RETURNING (id)",
		uC.Login,
		hashedPassword,
	)
	defer rows.Close()
	if err != nil {
		log.Errorf("error inserting new user: %s\n", err.Error())
		return &entity.User{}, err
	}

	newUser := uC.ToUserView(hashedPassword).ToEntity(oid.Undefined)
	if !rows.Next() {
		err := errors.New("error there's no returned id from sql")
		log.Errorln(err.Error())
		return newUser, err // TODO WHAT TO DO WTF???!!!?
	}
	var id string
	err = rows.Scan(&id)
	if err != nil {
		log.Errorf("error scanning new newUser id: %s\n", err.Error())
		id = string(oid.Undefined)
	}

	newUser.Id = oid.ToId(id)
	return newUser, nil
}

func (uRepo *userRepository) DeleteById(ctx context.Context, id oid.Id) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}
