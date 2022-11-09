package uploadbiz

import (
	"bytes"
	"context"
	"social-media-be/common"
	"social-media-be/components/cloudprovider"
	uploadmodel "social-media-be/modules/upload/model"
	"strings"
)

type UploadStore interface {
	Create(ctx context.Context, data *common.Image) error
	Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*common.Image, error)
}

type uploadBiz struct {
	provider cloudprovider.CloudProvider
	store    UploadStore
}

func NewUploadBiz(provider cloudprovider.CloudProvider, store UploadStore) *uploadBiz {
	return &uploadBiz{
		provider: provider,
		store:    store,
	}
}

func (biz *uploadBiz) Upload(
	ctx context.Context,
	data []byte,
	folder,
	fileName string,
) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := common.GetImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	img, err := biz.provider.UploadImage(ctx, data, fileName, folder)

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h

	biz.store.Create(ctx, img)

	return img, nil
}
