package model

import (
	"time"

	"github.com/uptrace/bun"
)

type QR struct {
	bun.BaseModel `bun:"qrcodes"`

	ID    int     `bun:"id,pk,nullzero"`
	Email *string `bun:"email,nullzero,notnull,unique,type:varchar(255)"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
