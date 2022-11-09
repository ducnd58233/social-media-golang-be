package cloudprovider

import (
	"context"
	"social-media-be/common"
)

type CloudProvider interface {
	// Upload image to Cloud and response URL
	UploadImage(
		ctx context.Context,
		data []byte,
		fileName,
		cloudFolder string,
	) (*common.Image, error)
}
