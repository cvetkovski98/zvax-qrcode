package delivery

import (
	"time"

	"github.com/cvetkovski98/zvax-common/gen/pbqr"
	"github.com/cvetkovski98/zvax-qrcode/internal/dto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateQRCodeRequestToDto(req *pbqr.CreateQRCodeRequest) *dto.CreateQRCode {
	return &dto.CreateQRCode{
		Email:   req.Email,
		Content: req.Content,
		Stored:  req.Stored,
	}
}

func GetQRCodeRequestToDto(req *pbqr.GetQRCodeRequest) *dto.GetQRCode {
	return &dto.GetQRCode{
		Email: req.Email,
	}
}

func timeToTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func QRCodeDtoToResponse(dto *dto.QR) *pbqr.QRCodeResponse {
	return &pbqr.QRCodeResponse{
		Qr:        dto.Content,
		Stored:    dto.Stored,
		Email:     dto.Email,
		Location:  dto.Location,
		CreatedAt: timeToTimestamp(dto.CreatedAt),
		UpdatedAt: timeToTimestamp(dto.UpdatedAt),
	}
}
