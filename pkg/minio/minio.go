package minio

import (
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
