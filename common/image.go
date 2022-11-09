package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io"
	"log"
)

type Image struct {
	Url       string `json:"url" gorm:"column:url;"`
	FileName  string `json:"file_name" gorm:"file_name;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
	Type      string `json:"type,omitempty" gorm:"-"`
}

func (Image) TableName() string { return "images" }

// Get data from database, json store in db is string
func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Faled to unmarshal JSONB value:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type Images []Image

func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *Images) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

func GetImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)

	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
