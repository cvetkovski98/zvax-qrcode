package dto

import "time"

type CreateQRCode struct {
	Email   *string
	Content string
	Stored  bool
}

type GetQRCode struct {
	Email string
}

type QR struct {
	Content   []byte
	Stored    bool
	Email     *string
	Location  *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
