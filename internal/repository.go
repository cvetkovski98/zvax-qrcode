package qrcode

import (
	"context"
	"net/url"

	"github.com/cvetkovski98/zvax-qrcode/internal/model"
)

type Repository interface {
	InsertOne(context.Context, *model.QR) (*model.QR, error)
	FindOneByEmail(context.Context, string) (*model.QR, error)
}

type ObjectStore interface {
	RemoveQR(ctx context.Context, email string) error
	UploadQR(ctx context.Context, email string, content []byte) (string, error)
	GetResourceLocation(ctx context.Context, objectName string) (*url.URL, error)
	CreateBucket(ctx context.Context) error
}
