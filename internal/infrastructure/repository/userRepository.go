package repository

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/internal/domain/storage/dto"
	"api-app/pkg/db/postgresql"
	"api-app/pkg/object/oid"
	"api-app/pkg/security/passlib"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type UserRepository storage.UserStorage

type userRepository struct {
	pool        postgresql.PGXPool
	hashManager passlib.HashManager
}

var _ UserRepository = (*userRepository)(nil)

func NewUserRepository(pool postgresql.PGXPool, hashManager passlib.HashManager) UserRepository {
	return &userRepository{pool: pool, hashManager: hashManager}
}

func (uRepo *userRepository) GetByLoginAndHashedPassword(
	ctx context.Context,
	login, hashedPassword string,
) (*entity.User, error) {
	sqlQuery := "SELECT * FROM public.users WHERE login=$1 AND hashed_password=$2"
	row := uRepo.pool.QueryRow(ctx, sqlQuery, login, hashedPassword)

	user := &entity.User{}
	err := row.Scan(user)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, storage.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uRepo *userRepository) GetById(ctx context.Context, id oid.Id) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uRepo *userRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uRepo *userRepository) GetAllInMap(ctx context.Context) (map[oid.Id]*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uRepo *userRepository) Store(ctx context.Context, uC dto.UserCreate) (*entity.User, error) {
	hashedPassword, err := uRepo.hashManager.SHA256WithSalt(uC.Password)
	if err != nil {
		return &entity.User{}, err
	}
	newUser := uC.ToEntity(oid.Undefined, hashedPassword)
	rows, err := uRepo.pool.Query(
		ctx,
		"INSERT INTO public.users (login, hashed_password) VALUES ($1, $2) RETURNING (id)",
		newUser.Login,
		newUser.HashedPassword,
	)
	defer rows.Close()
	if err != nil {
		log.Errorf("error inserting new user: %s\n", err.Error())
		return &entity.User{}, err
	}

	if !rows.Next() {
		err := errors.New("error there's no returned id from sql")
		log.Errorln(err.Error())
		// TODO implement it without using id, instead use login
		//_, err = uRepo.DeleteById(ctx, newUser.Id)
		//if err != nil {
		//	log.Errorf("I even couldn't delete this user from repository (TT): %s", err.Error())
		//}
		return newUser, err
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

func (uRepo *userRepository) DeleteById(ctx context.Context, id oid.Id) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
