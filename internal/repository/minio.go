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
}

func (s *minioObjectStore) GetResourceLocation(ctx context.Context, bucket string, email string) (*url.URL, error) {
	expiry := 7 * 24 * time.Hour
	objectName := getObjectName(email)
	return s.client.PresignedGetObject(ctx, bucket, objectName, expiry, nil)
}

func (s *minioObjectStore) RemoveQR(ctx context.Context, bucket string, objectName string) error {
	return s.client.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
}

func (s *minioObjectStore) UploadQR(ctx context.Context, bucket string, email string, content []byte) (string, error) {
	objectName := getObjectName(email)
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

func (s *minioObjectStore) CreateBucket(ctx context.Context, bucketName string) error {
	if err := s.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
		exists, errBucketExists := s.client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s", bucketName)
		} else {
			return err
		}
	}
	log.Printf("Successfully created %s", bucketName)
	if err := s.client.SetBucketPolicy(ctx, bucketName, `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::`+bucketName+`/*"]}]} `); err != nil {
		log.Printf("Failed to set bucket permissions")
		return err
	}
	log.Printf("Successfully set %s permissions", bucketName)
	return nil

}

func NewMinioObjectStore(client *minio.Client) qrcode.ObjectStore {
	return &minioObjectStore{client: client}
}
