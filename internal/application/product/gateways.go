package product

import (
	"github.com/achillescres/saina-api/internal/config"
	parser "github.com/achillescres/saina-api/internal/infrastructure/gateway/tais"
	"github.com/achillescres/saina-api/pkg/aws/s3"
)

type Gateways struct {
	Services   *Services
	TaisParser parser.TaisParser
	TaisOutput parser.TaisOutput
}

func NewGateways(services *Services, cfg config.TaisConfig, bucket s3.Bucket) *Gateways {
	taisParser := parser.NewTaisParser(services.TaisParserService, cfg)
	taisOutput := parser.NewTaisOutput(bucket, cfg)
	return &Gateways{Services: services, TaisParser: taisParser, TaisOutput: taisOutput}
}
