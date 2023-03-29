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

func (r *walletRepo) CreateWallet(ctx context.Context, balance int64) (domain.Wallet, error) {
	w := domain.Wallet{
		Balance: balance,
	}

	if err := r.tx(ctx).Create(&w).Error; err != nil {
		return domain.Wallet{}, err
	}

	return w, nil
}

func (r *walletRepo) GetWalletByID(ctx context.Context, id int64) (domain.Wallet, error) {
	var w domain.Wallet
	if err := r.tx(ctx).Take(&w, id).Error; err != nil {
		return domain.Wallet{}, err
	}

	return w, nil
}
