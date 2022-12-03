package service

import (
	"context"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/service/sto"
	storage2 "github.com/achillescres/saina-api/internal/domain/storage"
	dto2 "github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/object/oid"
	"github.com/achillescres/saina-api/pkg/security/ajwt"
	"github.com/achillescres/saina-api/pkg/security/passlib"
	log "github.com/sirupsen/logrus"
)

type AuthService interface {
	RegisterUser(ctx context.Context, regInput *sto.RegisterUserInput) (oid.Id, error)
	LoginUser(ctx context.Context, loginUserInput *sto.LoginUserInput) (*string, *string, error)
}

type authService struct {
	userStorage         storage2.UserStorage
	refreshTokenStorage storage2.RefreshTokenStorage
	hasher              passlib.HashManager
	jwtManager          ajwt.JWTManager
	cfg                 config.AuthConfig
}

var _ AuthService = (*authService)(nil)

func NewAuthService(
	userStorage storage2.UserStorage,
	refreshTokenStorage storage2.RefreshTokenStorage,
	hasher passlib.HashManager,
	jwtManager ajwt.JWTManager,
	cfg config.AuthConfig,
) AuthService {
	return &authService{userStorage: userStorage, refreshTokenStorage: refreshTokenStorage, hasher: hasher, jwtManager: jwtManager, cfg: cfg}
}

func (aS *authService) RegisterUser(ctx context.Context, regUserInput *sto.RegisterUserInput) (oid.Id, error) {
	createUser := dto2.NewUserCreate(regUserInput.Login, regUserInput.Password)
	log.Infof("registering user login=%s\n", regUserInput.Login)
	user, err := aS.userStorage.Store(ctx, *createUser)
	id := user.Id
	if err != nil {
		return oid.Undefined, err
	}

	return id, nil
}

func (aS *authService) createRefreshToken(ctx context.Context) (*entity.RefreshToken, error) {
	token, createTime, expireTime, err := aS.jwtManager.RefreshToken()
	if err != nil {
		return nil, err
	}

	rTC := dto2.NewRefreshTokenCreate(token, createTime, expireTime)
	rT, err := aS.refreshTokenStorage.Store(ctx, *rTC)
	if err != nil {
		return nil, err
	}

	return rT, err
}

func (aS *authService) LoginUser(ctx context.Context, loginUserInput *sto.LoginUserInput) (jwt *string, rt *string, err error) {
	hashedPassword, err := aS.hasher.SHA256WithSalt(loginUserInput.Password)
	if err != nil {
		return nil, nil, err
	}

	user, err := aS.userStorage.GetByLoginAndHashedPassword(ctx, loginUserInput.Login, hashedPassword)
	if err != nil {
		return nil, nil, err
	}

	jwtT, err := aS.jwtManager.User(user.Login, user.HashedPassword)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := aS.createRefreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	return &jwtT, &refreshToken.Token, nil
}
