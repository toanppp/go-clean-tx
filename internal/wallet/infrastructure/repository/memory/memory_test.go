package memory

import (
	"context"
	"math/rand"
	"reflect"
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
