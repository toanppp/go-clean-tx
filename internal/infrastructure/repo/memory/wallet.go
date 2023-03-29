package memory

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/port"
)

type walletRepo struct {
	data      map[int64]domain.Wallet
	increment *int64
	mu        sync.Mutex
}

func NewWalletRepo(data map[int64]domain.Wallet, increment int64) port.WalletRepo {
	return &walletRepo{
		data:      data,
		increment: &increment,
	}
}

func (r *walletRepo) CreateWallet(_ context.Context, balance int64) (wallet domain.Wallet, err error) {
	wallet.ID = atomic.AddInt64(r.increment, 1)
	wallet.Balance = balance
	r.data[wallet.ID] = wallet
	return
}

func (r *walletRepo) GetWalletByID(_ context.Context, id int64) (domain.Wallet, error) {
	if wallet, ok := r.data[id]; ok {
		return wallet, nil
	}
	return domain.Wallet{}, domain.ErrorNotFound
}

func (r *walletRepo) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return tFunc(ctx)
}
