package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	awsManager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	awsTypes "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io"
	"strings"
	"time"
)

type Bucket interface {
	GetLatestFileKey(ctx context.Context) (string, time.Time, error)
	DownloadFileByKey(ctx context.Context, key string) ([]byte, int64, error)
	UploadLargeFile(ctx context.Context, key string, file io.Reader) error
	UploadFile(ctx context.Context, key string, file io.Reader) error
}

type bucket struct {
	name                         string
	fileDownloadTL, fileUploadTL time.Duration

	client     *s3.Client
	uploader   *awsManager.Uploader
	downloader *awsManager.Downloader

	lastFilename string
}

func NewBucket(ctx context.Context, bucketName string, fileDownloadTL time.Duration, fileUploadTL time.Duration) (Bucket, error) {
	if bucketName == "" {
		return nil, fmt.Errorf("error bucketName cant be empty")
	}
	if fileDownloadTL.Seconds() < 1 && fileDownloadTL.Seconds() > 30 {
		return nil, fmt.Errorf("error fileDownloadTL must be in range 1 - 30 seconds")
	}

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == "ru-central1" {
			return aws.Endpoint{
				PartitionID:   "yc",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1-a",
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	// Подгружаем конфигрурацию из ~/.aws/*
	//awsConfig.LoadSharedConfigProfile()
	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		return nil, err
	}
	// Создаем клиента для доступа к хранилищу S3
	client := s3.NewFromConfig(cfg)
	//err = bucketExists(ctx, client, bucketName)
	//if err != nil {
	//	return nil, err
	//}

	var partMBytes int64 = 10
	uploader := awsManager.NewUploader(client, func(u *awsManager.Uploader) {
		u.PartSize = partMBytes * 1024 * 1024
	})

	downloader := awsManager.NewDownloader(client)

	return &bucket{
		name:           bucketName,
		fileDownloadTL: fileDownloadTL,
		fileUploadTL:   fileUploadTL,

		client:     client,
		uploader:   uploader,
		downloader: downloader,
	}, err
}

func (b *bucket) UploadFile(ctx context.Context, key string, file io.Reader) error {
	_, err := b.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(b.name),
		Key:    aws.String(key),
		Body:   file,
	})
	return err
}

func (b *bucket) UploadLargeFile(ctx context.Context, key string, file io.Reader) error {
	err := bucketExists(ctx, b.client, b.name)
	if err != nil {
		return err
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*b.fileUploadTL)
	defer cancel()
	_, err = b.uploader.Upload(ctxTimeout, &s3.PutObjectInput{
		Bucket: aws.String(b.name),
		Key:    aws.String(key),
		Body:   file,
	})
	return err
}

func (b *bucket) GetLatestFileKey(ctx context.Context) (string, time.Time, error) {
	// Get all files in bucket
	result, err := b.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(b.name),
	})

	if err != nil {
		return "", time.Time{}, err
	}
	var latestObj awsTypes.Object
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
	obj, err := b.GetFileMetadata(ctx, key)
	if err != nil {
		return nil, 0, err
	}

	buf := make([]byte, obj.ContentLength)
	wr := awsManager.NewWriteAtBuffer(buf)
	ctxDw, cancel := context.WithTimeout(ctx, b.fileDownloadTL)
	defer cancel()
	numBytes, err := b.downloader.Download(
		ctxDw,
		wr,
		&s3.GetObjectInput{
			Bucket: aws.String(b.name),
			Key:    aws.String(key),
		},
		func(downloader *awsManager.Downloader) {
			downloader.Concurrency = 5
		},
	)
	if err != nil {
		return nil, 0, err
	}
	return buf, numBytes, nil
}

func (b *bucket) GetFileMetadata(ctx context.Context, key string) (*s3.HeadObjectOutput, error) {
	object, err := b.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(b.name),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func bucketExists(ctx context.Context, client *s3.Client, name string) error {
	_, err := client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(name),
	})
	return err
}
