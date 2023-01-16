package repository

import (
	"context"
	"fmt"

	qrcode "github.com/cvetkovski98/zvax-qrcode/internal"
	"github.com/cvetkovski98/zvax-qrcode/internal/config"
	"github.com/cvetkovski98/zvax-qrcode/pkg/gcp"
	"github.com/cvetkovski98/zvax-qrcode/pkg/minio"
)

func Create(ctx context.Context, cfg *config.Config) (qrcode.ObjectStore, error) {
	switch cfg.ObjectStore.Type {
	case "MinIO":
		client, err := minio.NewMinioClient(&cfg.ObjectStore)
		if err != nil {
			return nil, err
		}
		return NewMinioObjectStore(client), nil
	case "GCP":
		client, err := gcp.NewGCPClient(ctx)
		if err != nil {
			return nil, err
		}
		return NewGCPObjectStore(client), nil
	}
	return nil, fmt.Errorf("unsupported repository type: %s", cfg.ObjectStore.Type)
}
