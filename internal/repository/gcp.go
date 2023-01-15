package repository

import (
	"context"
	"net/url"

	qrcode "github.com/cvetkovski98/zvax-qrcode/internal"
	"cloud.google.com/go/storage"
)

type gcpObjectStore struct {
	client *storage.Client
}

// CreateBucket implements qrcode.ObjectStore
func (*gcpObjectStore) CreateBucket(context.Context, string) error {
	panic("unimplemented")
}

// GetResourceLocation implements qrcode.ObjectStore
func (*gcpObjectStore) GetResourceLocation(context.Context, string, string) (*url.URL, error) {
	panic("unimplemented")
}

// RemoveQR implements qrcode.ObjectStore
func (*gcpObjectStore) RemoveQR(context.Context, string, string) error {
	panic("unimplemented")
}

// UploadQR implements qrcode.ObjectStore
func (*gcpObjectStore) UploadQR(context.Context, string, string, []byte) (string, error) {
	panic("unimplemented")
}

func NewGCPObjectStore() qrcode.ObjectStore {
	return &gcpObjectStore{}
}
