package port

//go:generate mockgen -source=./wallet_repository.go -aux_files github.com/toanppp/go-clean-tx/internal/port=transactor.go -destination=./mock/mock_wallet_repository.go -package=mock

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/toanppp/go-clean-tx/internal/domain"
)

type WalletRepository interface {
	Transactor
	CreateWallet(ctx context.Context, balance int64) (wallet domain.Wallet, err error)
	GetWalletByID(ctx context.Context, id int64) (wallet domain.Wallet, err error)
}
