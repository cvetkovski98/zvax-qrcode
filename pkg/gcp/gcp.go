package gcp

import (
	"context"

	"cloud.google.com/go/storage"
)

func NewGCPClient(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx)
}
