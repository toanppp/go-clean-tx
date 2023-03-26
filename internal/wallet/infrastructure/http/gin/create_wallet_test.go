package gin

import (
	"errors"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/mock"
	"github.com/toanppp/go-clean-tx/internal/port"
	"github.com/toanppp/go-clean-tx/internal/wallet/infrastructure/http/presenter"
	"github.com/toanppp/go-clean-tx/pkg/assert"
	"github.com/toanppp/go-clean-tx/pkg/httpjson"
	"github.com/toanppp/go-clean-tx/pkg/response"
)

func newMockEngine(mockWalletUseCase port.WalletUseCase) *gin.Engine {
	mockEngine := gin.New()
	RegisterWalletHTTP(mockEngine.Group(""), mockWalletUseCase)
	return mockEngine
}

func TestCreateWallet(t *testing.T) {
	// mock use case
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	wallet := domain.Wallet{
		ID:      rand.Int63(),
		Balance: rand.Int63(),
	}
	mockWalletUseCase := mock.NewMockWalletUseCase(mockCtrl)
	mockWalletUseCase.EXPECT().CreateWallet(gomock.Any(), wallet.Balance).Return(wallet, nil)

	// mock router
	mockEngine := newMockEngine(mockWalletUseCase)

	// mock request
	mockBody := presenter.CreateWallet{
		Balance: wallet.Balance,
	}
	mockReq, err := httpjson.NewRequest(http.MethodPost, "/v1", mockBody)
	if err != nil {
		t.Fatalf("cannot create new httpjson request: %v", err)
	}

	// recorder response
	w := httptest.NewRecorder()
	mockEngine.ServeHTTP(w, mockReq)

	assert.StatusCode(t, w.Code, http.StatusOK)
	assert.JSON(t, w.Body.String(), response.Response[domain.Wallet]{
		Data:    wallet,
		Message: http.StatusText(http.StatusOK),
	})
}

func TestCreateWalletFail(t *testing.T) {
	// mock use case
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockWalletUseCase := mock.NewMockWalletUseCase(mockCtrl)

	// mock router
	mockEngine := newMockEngine(mockWalletUseCase)

	tests := []struct {
		name    string
		balance int64
	}{
		{
			"empty",
			0,
		},
		{
			"negative",
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock request
			body := presenter.CreateWallet{
				Balance: tt.balance,
			}
			mockReq, err := httpjson.NewRequest(http.MethodPost, "/v1", body)
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

func TestCreateWalletError(t *testing.T) {
	// mock use case
	mockErr := errors.New("cannot connect to the database")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockWalletUseCase := mock.NewMockWalletUseCase(mockCtrl)
	mockWalletUseCase.EXPECT().CreateWallet(gomock.Any(), gomock.Any()).Return(domain.Wallet{}, mockErr)

	// mock router
	mockEngine := newMockEngine(mockWalletUseCase)

	// mock request
	body := presenter.CreateWallet{
		Balance: rand.Int63(),
	}
	req, err := httpjson.NewRequest(http.MethodPost, "/v1", body)
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
