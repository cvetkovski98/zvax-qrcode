package mapper

import (
	"net/url"

	"github.com/cvetkovski98/zvax-qrcode/internal/dto"
	"github.com/cvetkovski98/zvax-qrcode/internal/model"
)

func QRDtoFromContent(content []byte) *dto.QR {
	return &dto.QR{
		Content: content,
		Stored:  false,
	}
}

func QRDtoFromModel(qr *model.QR, location *url.URL) *dto.QR {
	location.RawQuery = ""
	l := location.String()
	return &dto.QR{
		Content:   nil,
		Stored:    true,
		Location:  &l,
		Email:     qr.Email,
		CreatedAt: &qr.CreatedAt,
		UpdatedAt: &qr.UpdatedAt,
	}
}
