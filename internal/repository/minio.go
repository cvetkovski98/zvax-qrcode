package repository

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"net/url"
	"time"

	qrcode "github.com/cvetkovski98/zvax-qrcode/internal"
	"github.com/minio/minio-go/v7"
)

type minioObjectStore struct {
	client *minio.Client
}

func (*minioObjectStore) getObjectName(email string) string {
	emailBytes := []byte(email)
	emailHash := sha256.Sum256(emailBytes)
	return fmt.Sprintf("%x.png", emailHash)
}

func (s *minioObjectStore) GetResourceLocation(ctx context.Context, bucket string, email string) (*url.URL, error) {
	expiry := 7 * 24 * time.Hour
	objectName := s.getObjectName(email)
	return s.client.PresignedGetObject(ctx, bucket, objectName, expiry, nil)
}

func (s *minioObjectStore) RemoveQR(ctx context.Context, bucket string, email string) error {
	objectName := s.getObjectName(email)
	return s.client.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
}

func (s *minioObjectStore) UploadQR(ctx context.Context, bucket string, email string, content []byte) (string, error) {
	objectName := s.getObjectName(email)
	contentReader := bytes.NewReader(content)
	contentLenght := int64(len(content))
	options := minio.PutObjectOptions{
		ContentType: "image/png",
	}
	if _, err := s.client.PutObject(ctx, bucket, objectName, contentReader, contentLenght, options); err != nil {
		return "", err
	}
	return objectName, nil
}

func NewMinioObjectStore(client *minio.Client) qrcode.ObjectStore {
	return &minioObjectStore{client: client}
}
