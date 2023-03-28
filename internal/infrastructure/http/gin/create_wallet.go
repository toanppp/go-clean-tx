package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/infrastructure/http/presenter"
)

func (h *walletHandler) createWallet(ctx *gin.Context) {
	var req presenter.CreateWallet
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responseBadRequest(ctx, err.Error())
		return
	}

	w, err := h.walletUseCase.CreateWallet(ctx, req.Balance)
	if err != nil {
		responseError(ctx, err)
		return
	}

	responseSuccess(ctx, w)
}
