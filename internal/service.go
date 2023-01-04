package qrcode

import (
	"context"

	"github.com/cvetkovski98/zvax-qrcode/internal/dto"
)

type Service interface {
	CreateQRCode(context.Context, *dto.CreateQRCode) (*dto.QR, error)
	GetQRCode(context.Context, *dto.GetQRCode) (*dto.QR, error)
}
