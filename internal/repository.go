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
	RemoveQR(context.Context, string, string) error
	UploadQR(context.Context, string, string, []byte) (string, error)
	GetResourceLocation(context.Context, string, string) (*url.URL, error)
	CreateBucket(context.Context, string) error
}
