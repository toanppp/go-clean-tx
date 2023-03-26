package wallet_use_case

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
	w, err := u.walletRepository.CreateWallet(ctx, balance)
	if err != nil {
		return domain.Wallet{}, err
	}

	return w, nil
}

func (u *walletUseCase) GetBalanceByID(ctx context.Context, id int64) (int64, error) {
	w, err := u.walletRepository.GetWalletByID(ctx, id)
	if err != nil {
		return 0, err
	}

	return w.Balance, nil
}
