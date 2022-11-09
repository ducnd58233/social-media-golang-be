package awsprovider

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
)

type S3 interface {
	// UploadFileData(ctx context.Context, data []byte, fileName string) (string, error)
	// Upload image to AWS S3 and response URL
	UploadImage(ctx context.Context, fileName string, cloudFolder string) (string, error)
	// Get image link from uploaded with imageKey and duration
	// GetImageWithExpireLink(ctx context.Context, imageKey string, duration time.Duration) (string, error)
	// // Delete image with imageKey and duration
	// DeleteImages(ctx context.Context, fileKeys []string) error
	// // Delete any object
	// DeleteObject(ctx context.Context, key string) error
}

type s3Provider struct {
	name    string
	session *session.Session
	logger  *logrus.Entry
	cfg     *s3Config
	service *s3.S3
}

type s3Config struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
}

func NewS3Provider(bucketName, region, apiKey, secret, domain string, logger *logrus.Logger) *s3Provider {
	provider := &s3Provider{
		name: "aws-s3",
		cfg: &s3Config{
			bucketName: bucketName,
			region:     region,
			apiKey:     apiKey,
			secret:     secret,
			domain:     domain,
		},
		logger: logger.WithField("service_name", "s3Provider"),
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.cfg.region),
		Credentials: credentials.NewStaticCredentials(
			provider.cfg.apiKey,
			provider.cfg.secret,
			"",
		),
	})

	if err != nil {
		provider.logger.Error(err)
		return nil
	}

	provider.session = s3Session
	provider.service = s3.New(provider.session)

	return provider
}
