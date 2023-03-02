// Package service  @Author xiaobaiio 2023/2/22 14:47:00
package service

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"os"
	"sync"
)

type ImageStore interface {
	Save(laptapId string, imageType string, imgData bytes.Buffer) (string, error)
}
type DiskImageStore struct {
	mutex       sync.RWMutex
	imageFolder string
	images      map[string]*ImageInfo
}

type ImageInfo struct {
	LaptapId string
	Type     string
	Path     string
}

func NewDiskImageStore(imageFolder string) *DiskImageStore {
	return &DiskImageStore{
		images:      make(map[string]*ImageInfo),
		imageFolder: imageFolder,
	}
}
func (store *DiskImageStore) Save(laptapId string, imageType string, imgData bytes.Buffer) (string, error) {
	imageId, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}
	imagePath := fmt.Sprintf("%s/%s%s", store.imageFolder, imageId, imageType)
	file, err := os.Create(imagePath)
	defer file.Close()
	if err != nil {
		return "", fmt.Errorf("cannot create img file: %w", err)
	}
	_, err = imgData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image data to file: %w", err)
	}
	store.mutex.Lock()
	defer store.mutex.Unlock()
	imageInfo := &ImageInfo{
		LaptapId: laptapId,
		Type:     imageType,
		Path:     imagePath,
	}
	store.images[imageId.String()] = imageInfo
	return imageId.String(), nil
}
