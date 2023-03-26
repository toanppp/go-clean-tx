package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/wallet/infrastructure/http/presenter"
	"github.com/toanppp/go-clean-tx/pkg/ginresp"
)

func (h *walletHandler) createWallet(ctx *gin.Context) {
	var req presenter.CreateWallet
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ginresp.ResponseBadRequest(ctx, err.Error())
		return
	}

	w, err := h.walletUseCase.CreateWallet(ctx, req.Balance)
	if err != nil {
		ginresp.ResponseError(ctx, err)
		return
	}

	ginresp.ResponseSuccess(ctx, w)
}
