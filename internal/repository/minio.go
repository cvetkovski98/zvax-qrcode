package repository

import (
	"bytes"
	"context"
	"log"
	"net/url"
	"time"

	qrcode "github.com/cvetkovski98/zvax-qrcode/internal"
	"github.com/minio/minio-go/v7"
)

type minioObjectStore struct {
	client *minio.Client
	bucket string
}

func (s *minioObjectStore) GetResourceLocation(ctx context.Context, objectName string) (*url.URL, error) {
	expiry := 7 * 24 * time.Hour
	return s.client.PresignedGetObject(ctx, s.bucket, objectName, expiry, nil)
}

func (s *minioObjectStore) RemoveQR(ctx context.Context, objectName string) error {
	return s.client.RemoveObject(ctx, s.bucket, objectName, minio.RemoveObjectOptions{})
}

func (s *minioObjectStore) UploadQR(ctx context.Context, email string, content []byte) (string, error) {
	objectName := getObjectName(email)
	contentReader := bytes.NewReader(content)
	contentLenght := int64(len(content))
	options := minio.PutObjectOptions{
		ContentType: "image/png",
	}
	if _, err := s.client.PutObject(ctx, s.bucket, objectName, contentReader, contentLenght, options); err != nil {
		return "", err
	}
	return objectName, nil
}

func (s *minioObjectStore) CreateBucket(ctx context.Context) error {
	if err := s.client.MakeBucket(ctx, s.bucket, minio.MakeBucketOptions{}); err != nil {
		exists, errBucketExists := s.client.BucketExists(ctx, s.bucket)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s", s.bucket)
		} else {
			return err
		}
	}
	log.Printf("Successfully created %s", s.bucket)
	if err := s.client.SetBucketPolicy(ctx, s.bucket, `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::`+s.bucket+`/*"]}]} `); err != nil {
		log.Printf("Failed to set bucket permissions")
		return err
	}
	log.Printf("Successfully set %s permissions", s.bucket)
	return nil

}

func NewMinioObjectStore(client *minio.Client, bucket string) qrcode.ObjectStore {
	return &minioObjectStore{
		client: client,
		bucket: bucket,
	}
}
