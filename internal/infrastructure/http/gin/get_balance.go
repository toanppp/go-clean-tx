package gin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/infrastructure/http/presenter"
)

func (h *walletHandler) getBalance(ctx *gin.Context) {
	query, ok := ctx.GetQuery("id")
	if !ok {
		responseBadRequest(ctx, "require id")
		return
	}

	id, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		responseBadRequest(ctx, "invalid id")
		return
	}

	if id <= 0 {
		responseBadRequest(ctx, "id must be greater than 0")
		return
	}

	b, err := h.walletUseCase.GetBalanceByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			responseFail(ctx, http.StatusNotFound, "wallet not found", err)
			return
		}

		responseError(ctx, err)
		return
	}

	responseSuccess(ctx, presenter.Balance{
		Balance: b,
	})
}
