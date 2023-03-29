package memory

import (
	"context"
	"math/rand"
	"reflect"
	"sync"
	"testing"

	"github.com/toanppp/go-clean-tx/internal/domain"
)

func TestWalletRepository_CreateWallet(t *testing.T) {
	increment := rand.Int63n(100)
	repo := NewWalletRepository(map[int64]domain.Wallet{}, increment)

	b1 := rand.Int63()
	w1, err := repo.CreateWallet(context.Background(), b1)

	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	if w1.ID == 0 {
		t.Errorf("invalid generated id: got %d", w1.ID)
	}

	if w1.Balance != b1 {
		t.Errorf("wrong balance: got %d: want %d", w1.Balance, 5)
	}

	b2 := rand.Int63()
	w2, err := repo.CreateWallet(context.Background(), b2)

	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	if w2.Balance != b2 {
		t.Errorf("wrong balance: got %d: want %d", w2.Balance, 10)
	}

	if w1.ID >= w2.ID {
		t.Errorf("increment id error: %d - %d", w1.ID, w2.ID)
	}
}

func TestWalletRepository_WithinTransaction(t *testing.T) {
	repo := NewWalletRepository(map[int64]domain.Wallet{}, 0)
	mapID := make(map[int64]any, 100)

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			var w domain.Wallet
			b := rand.Int63()

			err := repo.WithinTransaction(context.Background(), func(ctx context.Context) (err error) {
				w, err = repo.CreateWallet(context.Background(), b)
				return
			})

			if err != nil {
				t.Errorf("an error occurred: %v", err)
				return
			}

			_, ok := mapID[w.ID]
			if ok {
				t.Errorf("duplicate id: %d", w.ID)
			}

			if w.Balance != b {
				t.Errorf("wrong balance: got %d: want %d", w.Balance, 5)
			}
		}()
	}

	wg.Wait()
}

func TestWalletRepository_GetWalletByID(t *testing.T) {
	want := domain.Wallet{
		ID:      rand.Int63(),
		Balance: rand.Int63(),
	}

	repo := NewWalletRepository(map[int64]domain.Wallet{
		want.ID: want,
	}, rand.Int63())

	got, err := repo.GetWalletByID(context.Background(), want.ID)
	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}
