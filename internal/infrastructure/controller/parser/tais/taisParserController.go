package parser

import (
	"context"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/pkg/aws/s3"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

type TaisParserController interface {
	RunFTPWatcher(ctx context.Context, hereIsNewFile chan<- string) error
}

type taisParserController struct {
	cfg          config.TaisParserControllerConfig
	taisParser   TaisParser
	bucket       s3.Bucket
	lastFileKey  string
	lastFileDate time.Time
}

func (taisParserC *taisParserController) RunFTPWatcher(ctx context.Context, hereIsNewFile chan<- string) error {
	t := time.NewTicker(taisParserC.cfg.FTPCheckTimeout)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			key, date, err := taisParserC.bucket.GetLatestFileKey(ctx)
			if err != nil {
				log.Errorf("error gettingLast file from bucket: %s", err)
				continue
			}

			if key == taisParserC.lastFileKey || date.Unix() <= taisParserC.lastFileDate.Unix() {
				continue
			}

			bytes, numBytes, err := taisParserC.bucket.DownloadFileByKey(ctx, key)
			if err != nil {
				log.Errorf("error couldn't download file from bucket: %s", err)
				continue
			}
			if bytes == nil || numBytes == 0 {
				log.Errorf("error downloaded file is empty")
				continue
			}

			log.Infof("file %s with size of %d bytes downloaded", key, numBytes)
			err = os.WriteFile(path.Join(taisParserC.cfg.TaisDirAbsPath, key), bytes, 0644)

			if err != nil {
				log.Errorf("error couldn't write file: %s", err)
				continue
			}

			hereIsNewFile <- key
		}
	}
}

func NewTaisParserController(ctx context.Context, cfg config.TaisParserControllerConfig, taisParser TaisParser) (TaisParserController, error) {
	bucket, err := s3.NewBucket(ctx, cfg.BucketName, cfg.FileDownloadTimeLimit)
	if err != nil {
		return nil, err
	}

	return &taisParserController{
		cfg:          cfg,
		taisParser:   taisParser,
		bucket:       bucket,
		lastFileKey:  "",
		lastFileDate: time.Date(2004, 12, 03, 1, 1, 1, 1, time.UTC),
	}, nil
}
