package port

//go:generate mockgen -source=./wallet_repository.go -destination=../mock/mock_wallet_repository.go -package=mock

import (
	"context"

	"github.com/toanppp/go-clean-tx/internal/domain"
)

type WalletRepository interface {
	CreateWallet(ctx context.Context, balance int64) (wallet domain.Wallet, err error)
	GetWalletByID(ctx context.Context, id int64) (wallet domain.Wallet, err error)
}
