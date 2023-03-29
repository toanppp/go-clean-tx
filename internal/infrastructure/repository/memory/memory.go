package memory

import (
	"context"
	"sync"

	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/port"
)

type walletRepository struct {
	data      map[int64]domain.Wallet
	increment int64
	mu        sync.Mutex
}

func NewWalletRepository(data map[int64]domain.Wallet, increment int64) port.WalletRepository {
	return &walletRepository{
		data:      data,
		increment: increment,
	}
}

func (r *walletRepository) CreateWallet(_ context.Context, balance int64) (domain.Wallet, error) {
	r.increment++

	w := domain.Wallet{
		ID:      r.increment,
		Balance: balance,
	}

	r.data[w.ID] = w
	return w, nil
}

func (r *walletRepository) GetWalletByID(_ context.Context, id int64) (domain.Wallet, error) {
	w, ok := r.data[id]
	if !ok {
		return domain.Wallet{}, domain.ErrorNotFound
	}

	return w, nil
}

func (r *walletRepository) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return tFunc(ctx)
}
