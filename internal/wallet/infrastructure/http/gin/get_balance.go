package gin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/wallet/infrastructure/http/presenter"
	"github.com/toanppp/go-clean-tx/pkg/ginresp"
)

func (h *walletHandler) getBalance(ctx *gin.Context) {
	query, ok := ctx.GetQuery("id")
	if !ok {
		ginresp.ResponseBadRequest(ctx, "require id")
		return
	}

	id, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		ginresp.ResponseBadRequest(ctx, "invalid id")
		return
	}

	if id <= 0 {
		ginresp.ResponseBadRequest(ctx, "id must be greater than 0")
		return
	}

	b, err := h.walletUseCase.GetBalanceByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			ginresp.ResponseFail(ctx, http.StatusNotFound, "wallet not found", err)
			return
		}

		ginresp.ResponseError(ctx, err)
		return
	}

	ginresp.ResponseSuccess(ctx, presenter.Balance{
		Balance: b,
	})
}
