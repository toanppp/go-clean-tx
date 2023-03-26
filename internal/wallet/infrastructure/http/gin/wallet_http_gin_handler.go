package wallet_http_gin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/port"
	"github.com/toanppp/go-clean-tx/internal/wallet/infrastructure/http/presenter"
	wallet_http_request "github.com/toanppp/go-clean-tx/internal/wallet/infrastructure/http/request"
	"github.com/toanppp/go-clean-tx/pkg/gin_response"
)

type walletHandler struct {
	walletUseCase port.WalletUseCase
}

func (h *walletHandler) createWallet(ctx *gin.Context) {
	var req wallet_http_request.CreateWallet
	if err := ctx.ShouldBindJSON(&req); err != nil {
		gin_response.ResponseBadRequest(ctx, err.Error())
		return
	}

	w, err := h.walletUseCase.CreateWallet(ctx, req.Balance)
	if err != nil {
		gin_response.ResponseError(ctx, err)
		return
	}

	gin_response.ResponseSuccess(ctx, w)
}

func (h *walletHandler) getBalance(ctx *gin.Context) {
	query, ok := ctx.GetQuery("id")
	if !ok {
		gin_response.ResponseBadRequest(ctx, "require id")
		return
	}

	id, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		gin_response.ResponseBadRequest(ctx, "invalid id")
		return
	}

	if id <= 0 {
		gin_response.ResponseBadRequest(ctx, "id must be greater than 0")
		return
	}

	b, err := h.walletUseCase.GetBalanceByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			gin_response.ResponseFail(ctx, http.StatusNotFound, "wallet not found", err)
			return
		}

		gin_response.ResponseError(ctx, err)
		return
	}

	gin_response.ResponseSuccess(ctx, wallet_http_presenter.Balance{
		Balance: b,
	})
}
