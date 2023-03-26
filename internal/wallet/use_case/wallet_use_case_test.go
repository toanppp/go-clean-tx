package wallet_use_case

import (
	"context"
	"math/rand"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/mock"
)

func TestWalletUseCase_CreateWallet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockWalletRepository(mockCtrl)
	testUseCase := NewWalletUseCase(mockRepo)
	ctx := context.Background()
	wallet := domain.Wallet{
		ID:      rand.Int63(),
		Balance: rand.Int63(),
	}

	mockRepo.EXPECT().CreateWallet(ctx, wallet.Balance).Return(wallet, nil)

	got, err := testUseCase.CreateWallet(context.Background(), wallet.Balance)

	if err != nil {
		t.Fatalf("an error occurred: %v", err)
	}

	if !reflect.DeepEqual(got, wallet) {
		t.Errorf("wrong balance: got :%v, want: %v", got, wallet)
	}
}

func TestWalletUseCase_GetBalanceByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockWalletRepository(mockCtrl)
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
