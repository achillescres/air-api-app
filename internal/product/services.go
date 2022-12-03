package product

import (
	config2 "github.com/achillescres/saina-api/internal/config"
	service2 "github.com/achillescres/saina-api/internal/domain/service"
	"github.com/achillescres/saina-api/pkg/security/ajwt"
	"github.com/achillescres/saina-api/pkg/security/passlib"
)

type Services struct {
	AuthService   service2.AuthService
	ParserService service2.ParserService
	TablesService service2.DataService
}

func NewServices(
	repos *Repositories,
	taisParserConfig *config2.TaisParserConfig,
	hasher passlib.HashManager,
	jwtManager ajwt.JWTManager,
	cfg config2.AuthConfig,
) (*Services, error) {
	return &Services{
		AuthService:   service2.NewAuthService(repos.UserRepo, repos.RefreshTokenRepo, hasher, jwtManager, cfg),
		ParserService: service2.NewParserService(repos.FlightRepo, repos.TicketRepo, taisParserConfig),
		TablesService: service2.NewDataService(repos.FlightRepo, repos.TicketRepo),
	}, nil
}
