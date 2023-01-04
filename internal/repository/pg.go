package repository

import (
	"context"

	qrcode "github.com/cvetkovski98/zvax-qrcode/internal"
	"github.com/cvetkovski98/zvax-qrcode/internal/model"
	"github.com/uptrace/bun"
)

type pg struct {
	db *bun.DB
}

func (repository *pg) FindOneByEmail(ctx context.Context, email string) (*model.QR, error) {
	qr := new(model.QR)
	query := repository.db.NewSelect().Model(qr).Where("email = ?", email)
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}
	return qr, nil
}

func (repository *pg) InsertOne(ctx context.Context, qr *model.QR) (*model.QR, error) {
	if _, err := repository.db.NewInsert().Model(qr).Exec(ctx); err != nil {
		return nil, err
	}
	return qr, nil
}

func NewPgQRCodeRepository(db *bun.DB) qrcode.Repository {
	return &pg{db: db}
}
