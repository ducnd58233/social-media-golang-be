package cloudinaryprovider

import (
	"context"

	cld "github.com/cloudinary/cloudinary-go/v2"
	"github.com/sirupsen/logrus"
)

type Cloudinary interface {
	UploadImage(ctx context.Context, fileName string, cloudFolder string) (string, error)
}

type cloudinaryProvider struct {
	name    string
	logger  *logrus.Entry
	cfg     *cloudinaryConfig
	service *cld.Cloudinary
}

type cloudinaryConfig struct {
	apiKey    string
	secret    string
	cloudName string
}

func NewCloudinaryProvider(cloudName, apiKey, secret string, logger *logrus.Logger) *cloudinaryProvider {
	provider := &cloudinaryProvider{
		name: "cloudinary",
		cfg: &cloudinaryConfig{
			apiKey:    apiKey,
			secret:    secret,
			cloudName: cloudName,
		},
		logger: logger.WithField("service_name", "cloudinaryProvider"),
	}

	service, err := cld.NewFromParams(cloudName, apiKey, secret)

	if err != nil {
		provider.logger.Error(err)
		return nil
	}

	provider.service = service

	return provider
}

