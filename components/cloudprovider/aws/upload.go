package awsprovider

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"social-media-be/common"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (provider *s3Provider) Upload(ctx context.Context, fileName, cloudFolder string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)

	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	ext := filepath.Ext(file.Name())
	newFileName := fmt.Sprintf("/%s/%s", time.Now(), ext)

	fileKey := fmt.Sprintf("/%s/%s", cloudFolder, newFileName)
	params := &s3.PutObjectInput{
		Bucket:        aws.String(provider.cfg.bucketName),
		Key:           aws.String(fileKey), // filename
		ACL:           aws.String("private"),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
		Body:          fileBytes,
	}

	_, err = provider.service.PutObjectWithContext(ctx, params)
	if err != nil {
		return "", err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("https://%s.s3.amazonaws.com%s", provider.cfg.bucketName, fileKey),
		CloudName: "s3",
		Extension: ext,
	}

	return img.Url, nil
}

func (provider *s3Provider) UploadImage(
	ctx context.Context,
	data []byte,
	fileName,
	cloudFolder string,
) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	ext := filepath.Ext(fileName)

	newFileName := fmt.Sprintf("/%s/%s", time.Now(), ext)

	fileKey := fmt.Sprintf("/%s", newFileName)
	_, err := provider.service.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(provider.cfg.bucketName),
		Key:           aws.String(fileKey), // filename
		ACL:           aws.String("private"),
		ContentType:   aws.String(fileType),
		Body:          fileBytes,
	})

	if err != nil {
		return nil, err
	}

	img := &common.Image{
		FileName:  newFileName,
		Url:       fmt.Sprintf("https://%s.s3.amazonaws.com%s", provider.cfg.bucketName, fileKey),
		CloudName: "s3",
		Extension: ext,
		Type:      fileType,
	}

	return img, nil
}
