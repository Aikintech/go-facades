package gofacades

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type R2 struct {
	instance *s3.Client
	ctx      context.Context
	bucket   string
	url      string
}

func NewR2(config S3Config) *R2 {
	var (
		accessKeyId     = config.AccessKeyId
		accessKeySecret = config.AccessKeySecret
		bucket          = config.Bucket
		region          = config.Region
		accountId       = config.AccountId
		url             = config.CdnURL
	)
	if accessKeyId == "" || accessKeySecret == "" || url == nil || accountId == nil || bucket == "" || region == "" {
		log.Fatal("please set configuration first")
	}

	u := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", *accountId)

	fmt.Println(u)

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, opt ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: u}, nil
	})
	cfg, err := cfg.LoadDefaultConfig(context.TODO(),
		cfg.WithEndpointResolverWithOptions(r2Resolver),
		cfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		cfg.WithRegion("auto"),
	)

	if err != nil {
		log.Fatal(err)
	}

	instance := s3.NewFromConfig(cfg)

	return &R2{
		instance: instance,
		ctx:      context.Background(),
		bucket:   bucket,
		url:      *url,
	}
}

func (s *R2) Exist(path string) bool {
	_, err := s.instance.HeadObject(s.ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})

	return err == nil
}

func (s *R2) Delete(files ...string) error {
	var objectIdentifiers []s3Types.ObjectIdentifier
	for _, file := range files {
		objectIdentifiers = append(objectIdentifiers, s3Types.ObjectIdentifier{
			Key: aws.String(file),
		})
	}

	_, err := s.instance.DeleteObjects(s.ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String(s.bucket),
		Delete: &s3Types.Delete{
			Objects: objectIdentifiers,
			Quiet:   aws.Bool(true),
		},
	})

	return err
}

func (s *R2) GetBytes(file string) ([]byte, error) {
	resp, err := s.instance.GetObject(s.ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(file),
	})
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return data, nil
}

// func (s *R2) MakeDirectory(directory string) error {
// 	return nil
// }

func (s *R2) MimeType(file string) (string, error) {
	resp, err := s.instance.HeadObject(s.ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(file),
	})
	if err != nil {
		return "", err
	}

	return aws.ToString(resp.ContentType), nil
}

func (s *R2) Put(path string, source *multipart.FileHeader) (FileDetails, error) {
	return FileDetails{}, nil
}

func (s *R2) PutFile(path string, source *multipart.FileHeader) (FileDetails, error) {
	if source == nil {
		return FileDetails{}, fmt.Errorf("multipart.FileHeader is nil")
	}

	fileContents, err := source.Open()
	if err != nil {
		return FileDetails{}, err
	}
	defer fileContents.Close()

	// Get file details
	mimeType := source.Header.Get("Content-Type")
	extension := strings.ReplaceAll(filepath.Ext(source.Filename), ".", "")
	filename := fmt.Sprintf("%s.%s", GenerateRandomString(uint(20), true), extension)
	fullPath := filepath.Join(strings.ToLower(path), filename)

	uploader := manager.NewUploader(s.instance)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fullPath),
		Body:   fileContents,
	})
	fmt.Println(result)

	if err != nil {
		return FileDetails{}, err
	}

	return FileDetails{
		Disk:      "s3",
		Extension: extension,
		FileName:  filename,
		Size:      source.Size,
		Path:      fullPath,
		URL:       s.Url(fullPath),
		MimeType:  mimeType,
	}, nil
}

func (s *R2) Size(file string) (int64, error) {
	resp, err := s.instance.HeadObject(s.ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(file),
	})
	if err != nil {
		return 0, err
	}

	return *resp.ContentLength, nil
}

func (s *R2) Url(key string) string {
	if len(key) == 0 {
		return ""
	}

	return fmt.Sprintf("%s/%s", s.url, key)
}
