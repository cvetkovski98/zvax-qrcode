package service

import (
	"context"
	"errors"

	qrcode "github.com/cvetkovski98/zvax-qrcode/internal"
	"github.com/cvetkovski98/zvax-qrcode/internal/dto"
	"github.com/cvetkovski98/zvax-qrcode/internal/mapper"
	"github.com/cvetkovski98/zvax-qrcode/internal/model"
	errutil "github.com/cvetkovski98/zvax-qrcode/internal/utils/error"
	qrutil "github.com/cvetkovski98/zvax-qrcode/internal/utils/qr"
)

const bucket = "qrcodes"

type impl struct {
	r  qrcode.Repository
	os qrcode.ObjectStore
}

func (s *impl) validateCreateQRCode(ctx context.Context, request *dto.CreateQRCode) error {
	emailEmpty := request.Email == nil || *request.Email == ""

	if request.Stored && emailEmpty {
		return errors.New("email is required when storing QR code")
	}

	if !emailEmpty {
		qr, err := s.r.FindOneByEmail(ctx, *request.Email)
		err = errutil.ParseError(err)
		if err != nil {
			return err
		}
		if qr != nil {
			return errors.New("QR code already exists for this email")
		}
	}
	return nil
}

func (s *impl) CreateQRCode(ctx context.Context, request *dto.CreateQRCode) (*dto.QR, error) {
	if err := s.validateCreateQRCode(ctx, request); err != nil {
		return nil, err
	}
	content, err := qrutil.GenerateQRCode(request.Content)
	if err != nil {
		return nil, err
	}
	if !request.Stored {
		return mapper.QRDtoFromContent(content), nil
	}

	objectName, err := s.os.UploadQR(ctx, bucket, *request.Email, content)
	if err != nil {
		return nil, err
	}
	url, err := s.os.GetResourceLocation(ctx, bucket, objectName)
	if err != nil {
		go s.os.RemoveQR(ctx, bucket, *request.Email)
		return nil, err
	}
	QRin := &model.QR{
		Email: request.Email,
	}
	_, err = s.r.InsertOne(ctx, QRin)
	if err != nil {
		go s.os.RemoveQR(ctx, bucket, *request.Email)
		return nil, err
	}
	return mapper.QRDtoFromModel(QRin, url), nil
}

func (s *impl) GetQRCode(ctx context.Context, request *dto.GetQRCode) (*dto.QR, error) {
	result, err := s.r.FindOneByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	url, err := s.os.GetResourceLocation(ctx, bucket, request.Email)
	if err != nil {
		return nil, err
	}
	return mapper.QRDtoFromModel(result, url), nil
}

func NewQRCodeService(r qrcode.Repository, os qrcode.ObjectStore) qrcode.Service {
	return &impl{
		r:  r,
		os: os,
	}
}
