package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"strings"
	"time"
)

type Bucket interface {
	GetLatestFileKey(ctx context.Context) (string, time.Time, error)
	DownloadFileByKey(ctx context.Context, key string) ([]byte, int64, error)
}

type bucket struct {
	client                *s3.Client
	fileDownloadTimeLimit time.Duration
	name                  string

	lastFilename string
}

func NewBucket(ctx context.Context, bucketName string, fileDownloadTimeLimit time.Duration) (Bucket, error) {
	if bucketName == "" {
		return nil, fmt.Errorf("error NewBucket bucketName cant be empty")
	}

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == "ru-central1" {
			return aws.Endpoint{
				PartitionID:   "yc",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	// Подгружаем конфигрурацию из ~/.aws/*
	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		return nil, err
	}

	// Создаем клиента для доступа к хранилищу S3
	client := s3.NewFromConfig(cfg)
	return &bucket{
		client:                client,
		name:                  bucketName,
		fileDownloadTimeLimit: fileDownloadTimeLimit,
	}, err
}

func (b *bucket) GetLatestFileKey(ctx context.Context) (string, time.Time, error) {
	// Get all files in bucket
	result, err := b.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(b.name),
	})

	if err != nil {
		return "", time.Time{}, err
	}
	var latestObj types.Object
	var latestDate time.Time

	for i, c := range result.Contents {
		date, err := time.Parse("20220130", strings.TrimLeft(*c.Key, "REDATAFINAL"))
		if err != nil {
			date = *c.LastModified
		}
		if i == 0 || date.Unix() > latestDate.Unix() {
			latestDate = date
			latestObj = c
		}
	}

	return *latestObj.Key, latestDate, err
}

// DownloadFileByKey locks goroutine until the end of download or ctxDw or ctx
func (b *bucket) DownloadFileByKey(ctx context.Context, key string) ([]byte, int64, error) {
	dw := manager.NewDownloader(b.client)

	var buf []byte
	wr := manager.NewWriteAtBuffer(buf)
	ctxDw, cancel := context.WithTimeout(ctx, time.Second*120)
	defer cancel()
	numBytes, err := dw.Download(
		ctxDw,
		wr,
		&s3.GetObjectInput{
			Bucket: aws.String(b.name),
			Key:    aws.String(key),
		},
		func(downloader *manager.Downloader) {
			downloader.Concurrency = 5
		},
	)

	select {
	case <-ctx.Done():
		cancel()
		return nil, 0, nil
	case <-ctxDw.Done():
		if err != nil {
			return nil, 0, err
		}

		return buf, numBytes, nil
	}
}
