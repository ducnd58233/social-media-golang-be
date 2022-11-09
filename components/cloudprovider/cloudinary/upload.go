package cloudinaryprovider

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"social-media-be/common"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (c *cloudinaryProvider) UploadImage(
	ctx context.Context,
	data []byte,
	fileName,
	cloudFolder string,
) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	ext := filepath.Ext(fileName)
	newFileName := fmt.Sprintf("/%s/%s", time.Now(), ext)

	uploadResult, err := c.service.Upload.Upload(
		ctx,
		fileBytes,
		uploader.UploadParams{
			PublicID:    newFileName,
			AssetFolder: cloudFolder,
			Type:        api.DeliveryType(fileType),
			Async:       api.Bool(true),
		},
	)

	if err != nil {
		return nil, err
	}

	img := &common.Image{
		FileName:  newFileName,
		Url:       uploadResult.URL,
		CloudName: "cloudinary",
		Extension: ext,
		Type:      fileType,
	}

	return img, nil
}
