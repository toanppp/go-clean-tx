package usecase

import (
	"context"

	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/port"
)

type walletUseCase struct {
	walletRepository port.WalletRepository
}

func NewWalletUseCase(walletRepository port.WalletRepository) port.WalletUseCase {
	return &walletUseCase{
		walletRepository: walletRepository,
	}
}

func (u *walletUseCase) CreateWallet(ctx context.Context, balance int64) (domain.Wallet, error) {
	var wallet domain.Wallet

	err := u.walletRepository.WithinTransaction(ctx, func(txCtx context.Context) error {
		w, err := u.walletRepository.CreateWallet(txCtx, balance)
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
	w, err := u.walletRepository.GetWalletByID(ctx, id)
	if err != nil {
		return 0, err
	}

	return w.Balance, nil
}
