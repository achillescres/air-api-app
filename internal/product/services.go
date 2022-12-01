package product

import (
	"api-app/internal/config"
	"api-app/internal/domain/service"
	"api-app/pkg/security/ajwt"
	"api-app/pkg/security/passlib"
)

type Services struct {
	AuthService   service.AuthService
	ParserService service.ParserService
	TablesService service.DataService
}

func NewServices(
	repos *Repositories,
	taisParserConfig *config.TaisParserConfig,
	hasher passlib.HashManager,
	jwtManager ajwt.JWTManager,
	cfg config.AuthConfig,
) (*Services, error) {
	return &Services{
		AuthService:   service.NewAuthService(repos.UserRepo, repos.RefreshTokenRepo, hasher, jwtManager, cfg),
		ParserService: service.NewParserService(repos.FlightRepo, repos.TicketRepo, taisParserConfig),
		TablesService: service.NewDataService(repos.FlightRepo, repos.TicketRepo),
	}, nil
}
