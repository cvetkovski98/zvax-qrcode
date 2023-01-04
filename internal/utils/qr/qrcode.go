package qrutil

import (
	"bytes"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GenerateQRCode(content string) ([]byte, error) {
	QRCode, err := qr.Encode(content, qr.H, qr.Auto)
	if err != nil {
		return nil, err
	}
	QRCode, err = barcode.Scale(QRCode, 256, 256)
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	if err := png.Encode(&buffer, QRCode); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
