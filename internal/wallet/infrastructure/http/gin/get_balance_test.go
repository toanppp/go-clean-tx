package gin

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/mock"
	"github.com/toanppp/go-clean-tx/internal/wallet/infrastructure/http/presenter"
	"github.com/toanppp/go-clean-tx/pkg/assert"
	"github.com/toanppp/go-clean-tx/pkg/httpjson"
	"github.com/toanppp/go-clean-tx/pkg/response"
)

func TestGetBalance(t *testing.T) {
	// mock use case
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	wallet := domain.Wallet{
		ID:      rand.Int63(),
		Balance: rand.Int63(),
	}
	mockWalletUseCase := mock.NewMockWalletUseCase(mockCtrl)
	mockWalletUseCase.EXPECT().GetBalanceByID(gomock.Any(), wallet.ID).Return(wallet.Balance, nil)

	// mock router
	mockEngine := newMockEngine(mockWalletUseCase)

	// mock request
	mockReq, err := httpjson.NewRequest(http.MethodGet, fmt.Sprintf("/v1/balance?id=%d", wallet.ID), nil)
	if err != nil {
		t.Fatalf("cannot create new httpjson request: %v", err)
	}

	// recorder response
	w := httptest.NewRecorder()
	mockEngine.ServeHTTP(w, mockReq)

	assert.StatusCode(t, w.Code, http.StatusOK)
	assert.JSON(t, w.Body.String(), response.Response[presenter.Balance]{
		Data: presenter.Balance{
			Balance: wallet.Balance,
		},
		Message: http.StatusText(http.StatusOK),
	})
}

func TestGetBalanceFail(t *testing.T) {
	// mock use case
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockWalletUseCase := mock.NewMockWalletUseCase(mockCtrl)

	// mock router
	mockEngine := newMockEngine(mockWalletUseCase)

	tests := []struct {
		name   string
		params string
	}{
		{
			"empty",
			"",
		},
		{
			"empty string",
			"id=",
		},
		{
			"zero",
			"id=0",
		},
		{
			"negative",
			"id=-1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock request
			mockReq, err := httpjson.NewRequest(http.MethodGet, "/v1/balance?"+tt.params, nil)
			if err != nil {
				t.Fatalf("cannot create new httpjson request: %v", err)
			}

			// recorder response
			w := httptest.NewRecorder()
			mockEngine.ServeHTTP(w, mockReq)

			assert.StatusCode(t, w.Code, http.StatusBadRequest)
		})
	}

}

func TestGetBalanceError(t *testing.T) {
	// mock use case
	mockErr := errors.New("cannot connect to the database")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockWalletUseCase := mock.NewMockWalletUseCase(mockCtrl)
	mockWalletUseCase.EXPECT().GetBalanceByID(gomock.Any(), gomock.Any()).Return(int64(0), mockErr)

	// mock router
	mockEngine := newMockEngine(mockWalletUseCase)

	// mock request
	req, err := httpjson.NewRequest(http.MethodGet, fmt.Sprintf("/v1/balance?id=%d", rand.Int63()), nil)
	if err != nil {
		t.Fatalf("cannot create new httpjson request: %v", err)
	}

	// recorder response
	w := httptest.NewRecorder()
	mockEngine.ServeHTTP(w, req)

	assert.StatusCode(t, w.Code, http.StatusInternalServerError)
	assert.JSON(t, w.Body.String(), response.Response[any]{
		Data:    nil,
		Message: http.StatusText(http.StatusInternalServerError),
		Error:   mockErr.Error(),
	})
}
