package parser

import (
	"context"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/domain/object"
	"github.com/achillescres/saina-api/pkg/aws/s3"
	"strings"
	"time"
)

type TaisOutput interface {
	SendOutputTais(ctx context.Context, filename string, changes *object.TaisChanges) error
}

type taisOutput struct {
	bucket  s3.Bucket
	cfg     config.TaisConfig
	changes map[time.Time]object.TaisChanges
}

func NewTaisOutput(bucket s3.Bucket, cfg config.TaisConfig) TaisOutput {
	return &taisOutput{bucket: bucket, cfg: cfg, changes: make(map[time.Time]object.TaisChanges)}
}

func (tO *taisOutput) SendOutputTais(ctx context.Context, filename string, changes *object.TaisChanges) error {
	fileKey := filename
	tO.changes[time.Now()] = *changes
	// TODO create data collecting mechanism
	dataReader := strings.NewReader(changes.String())
	err := tO.bucket.UploadFile(ctx, fileKey, dataReader)
	return err
}
