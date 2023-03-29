package usecase

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/port/mock"
)

func TestWalletUseCase_CreateWallet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockWalletRepo(mockCtrl)
	testUseCase := NewWalletUseCase(mockRepo)
	ctx := context.Background()
	wallet := domain.Wallet{
		ID:      rand.Int63(),
		Balance: rand.Int63(),
	}

	mockRepo.EXPECT().CreateWallet(ctx, wallet.Balance).Return(wallet, nil)
	mockRepo.EXPECT().WithinTransaction(ctx, gomock.Any()).
		DoAndReturn(func(ctx context.Context, tFunc func(ctx context.Context) error) error {
			return tFunc(ctx)
		})

	got, err := testUseCase.CreateWallet(ctx, wallet.Balance)

	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	if !reflect.DeepEqual(got, wallet) {
		t.Errorf("wrong balance: got :%v, want: %v", got, wallet)
	}
}

func TestWalletUseCase_CreateWallet_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockWalletRepo(mockCtrl)
	testUseCase := NewWalletUseCase(mockRepo)
	ctx := context.Background()
	err := errors.New(http.StatusText(http.StatusInternalServerError))

	mockRepo.EXPECT().CreateWallet(ctx, gomock.Any()).Return(domain.Wallet{}, err)
	mockRepo.EXPECT().WithinTransaction(ctx, gomock.Any()).
		DoAndReturn(func(ctx context.Context, tFunc func(ctx context.Context) error) error {
			return tFunc(ctx)
		})

	_, e := testUseCase.CreateWallet(ctx, rand.Int63())
	if e == nil {
		t.Fatalf("not occur error")
	}

	if !errors.Is(e, err) {
		t.Fatalf("unexpected error got: %v, want: %v", e, err)
	}
}

func TestWalletUseCase_GetBalanceByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockWalletRepo(mockCtrl)
	testUseCase := NewWalletUseCase(mockRepo)
	ctx := context.Background()
	wallet := domain.Wallet{
		ID:      rand.Int63(),
		Balance: rand.Int63(),
	}

	mockRepo.EXPECT().GetWalletByID(ctx, wallet.ID).Return(wallet, nil)

	got, err := testUseCase.GetBalanceByID(context.Background(), wallet.ID)

	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	if got != wallet.Balance {
		t.Errorf("wrong balance: got :%v, want: %v", got, wallet.Balance)
	}
}
