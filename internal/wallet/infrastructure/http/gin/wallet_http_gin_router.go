package wallet_http_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/port"
)

func RegisterWalletGinHTTP(group *gin.RouterGroup, walletUseCase port.WalletUseCase) {
	h := walletHandler{
		walletUseCase: walletUseCase,
	}

	v1 := group.Group("v1")
	{
		v1.POST("", h.createWallet)
		v1.GET("balance", h.getBalance)
	}
}
