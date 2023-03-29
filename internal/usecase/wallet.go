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

func (u *walletUseCase) CreateWallet(ctx context.Context, balance int64) (domain.Wallet, error) {
	var wallet domain.Wallet

	err := u.walletRepo.WithinTransaction(ctx, func(txCtx context.Context) error {
		w, err := u.walletRepo.CreateWallet(txCtx, balance)
		if err != nil {
			return err
		}

		wallet = w
		return nil
	})

	if err != nil {
		return domain.Wallet{}, err
	}

	return wallet, nil
}

func (u *walletUseCase) GetBalanceByID(ctx context.Context, id int64) (int64, error) {
	w, err := u.walletRepo.GetWalletByID(ctx, id)
	if err != nil {
		return 0, err
	}

	return w.Balance, nil
}
