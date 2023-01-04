package minio

import (
	"context"
	"log"

	"github.com/cvetkovski98/zvax-common/pkg/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient(cfg *config.MinIO) (*minio.Client, error) {
	return minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.RootUser, cfg.RootPassword, ""),
		Secure: cfg.UseSSL,
	})
}

func CreateBucket(ctx context.Context, client *minio.Client, bucketName string) error {
	if err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s", bucketName)
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created %s", bucketName)
	}
	return nil
}

func MakeReadOnly(ctx context.Context, client *minio.Client, bucketName string) error {
	return client.SetBucketPolicy(ctx, bucketName, `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::`+bucketName+`/*"]}]} `)
}
