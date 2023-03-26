package port

//go:generate mockgen -source=./wallet_use_case.go -destination=../../mock/mock_wallet_use_case.go -package=mock

import (
	"context"

	"github.com/toanppp/go-clean-tx/internal/domain"
)

type WalletUseCase interface {
	CreateWallet(ctx context.Context, balance int64) (wallet domain.Wallet, err error)
	GetBalanceByID(ctx context.Context, id int64) (balance int64, err error)
}
