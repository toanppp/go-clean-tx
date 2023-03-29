package database

import (
	"context"

	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/port"
	"gorm.io/gorm"
)

type walletRepo struct {
	transactor
}

func NewWalletRepo(db *gorm.DB) port.WalletRepo {
	return &walletRepo{
		transactor: transactor{
			db: db,
		},
	}
}

func (r *walletRepo) CreateWallet(ctx context.Context, balance int64) (wallet domain.Wallet, err error) {
	wallet.Balance = balance
	err = r.tx(ctx).Create(&wallet).Error
	return
}

func (r *walletRepo) GetWalletByID(ctx context.Context, id int64) (wallet domain.Wallet, err error) {
	err = r.tx(ctx).Take(&wallet, id).Error
	return
}
