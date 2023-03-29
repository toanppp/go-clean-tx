package usecase

import (
	"context"

	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/port"
)

type walletUseCase struct {
	walletRepo port.WalletRepo
}

func NewWalletUseCase(walletRepo port.WalletRepo) port.WalletUseCase {
	return &walletUseCase{
		walletRepo: walletRepo,
	}
}

func (u *walletUseCase) CreateWallet(ctx context.Context, balance int64) (wallet domain.Wallet, err error) {
	err = u.walletRepo.WithinTransaction(ctx, func(txCtx context.Context) (err error) {
		wallet, err = u.walletRepo.CreateWallet(txCtx, balance)
		return
	})
	return
}

func (u *walletUseCase) GetBalanceByID(ctx context.Context, id int64) (int64, error) {
	wallet, err := u.walletRepo.GetWalletByID(ctx, id)
	return wallet.Balance, err
}
