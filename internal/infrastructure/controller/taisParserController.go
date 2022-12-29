package controller

import (
	"context"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/infrastructure/gateway/tais"
	"github.com/achillescres/saina-api/pkg/aws/s3"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"time"
)

type TaisParserController interface {
	RunFTPWatcher(ctx context.Context) error
}

type taisParserController struct {
	cfg                       config.TaisParserControllerConfig
	taisParser                parser.TaisParser
	bucket                    s3.Bucket
	lastFileKey               string
	lastFileDate              time.Time
	newTaisFileDownloadedChan chan<- string
}

func (taisParserC *taisParserController) RunFTPWatcher(ctx context.Context) error {
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

			if !strings.HasPrefix(key, "taisInput_") {
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
			log.Infof("Downloaded file with name %s", key)
			if bytes == nil || numBytes == 0 {
				log.Errorf("error downloaded file is empty")
				continue
			}

			log.Infof("file %s with size of %d bytes downloaded", key, numBytes)
			filename := "taisInput_" + key
			err = os.WriteFile(path.Join(taisParserC.cfg.TaisDirAbsPath, filename), bytes, 0644)

			if err != nil {
				log.Errorf("error couldn't write file: %s", err)
				continue
			}

			taisParserC.lastFileKey = key
			taisParserC.lastFileDate = date
			//taisParserC.newTaisFileDownloadedChan <- key
			_, errs, err := taisParserC.taisParser.ParseFirstTaisFile(ctx, filename)
			if err != nil {
				log.Errorf("error parsing file with name %s: %s", filename, err)
			}
			if len(errs.Errs) != 0 {
				log.Errorf("error there're errors while parsing tais file with name %s:, \n%v", filename, errs)
			}
		}
	}
}

func NewTaisParserController(
	cfg config.TaisParserControllerConfig,
	taisParser parser.TaisParser,
	bucket s3.Bucket,
) (TaisParserController, error) {
	return &taisParserController{
		cfg:                       cfg,
		taisParser:                taisParser,
		bucket:                    bucket,
		lastFileKey:               "",
		lastFileDate:              time.Date(2000, 01, 01, 1, 1, 1, 1, time.UTC),
		newTaisFileDownloadedChan: make(chan string, 1),
	}, nil
}
