package delivery

import (
	context "context"

	"github.com/cvetkovski98/zvax-common/gen/pbqr"
	qrcode "github.com/cvetkovski98/zvax-qrcode/internal"
)

type server struct {
	s qrcode.Service

	pbqr.UnimplementedQRCodeServer
}

func (s *server) CreateQRCode(ctx context.Context, request *pbqr.CreateQRCodeRequest) (*pbqr.QRCodeResponse, error) {
	in := CreateQRCodeRequestToDto(request)
	out, err := s.s.CreateQRCode(ctx, in)
	if err != nil {
		return nil, err
	}
	return QRCodeDtoToResponse(out), nil
}

func (s *server) GetQRCode(ctx context.Context, request *pbqr.GetQRCodeRequest) (*pbqr.QRCodeResponse, error) {
	in := GetQRCodeRequestToDto(request)
	out, err := s.s.GetQRCode(ctx, in)
	if err != nil {
		return nil, err
	}
	return QRCodeDtoToResponse(out), nil
}

func NewQRCodeServer(s qrcode.Service) pbqr.QRCodeServer {
	return &server{
		s: s,
	}
}
