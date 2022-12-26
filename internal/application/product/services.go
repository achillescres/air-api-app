package product

import (
	"github.com/achillescres/saina-api/internal/config"
	service "github.com/achillescres/saina-api/internal/domain/service"
	"github.com/achillescres/saina-api/pkg/security/ajwt"
	"github.com/achillescres/saina-api/pkg/security/passlib"
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
