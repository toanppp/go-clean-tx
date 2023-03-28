package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/port"
)

func RegisterHandle(engine *gin.Engine, walletUseCase port.WalletUseCase) {
	h := walletHandler{
		walletUseCase: walletUseCase,
	}

	v1 := engine.Group("v1")
	{
		wallet := v1.Group("wallet")
		{
			wallet.POST("", h.createWallet)
			wallet.GET("balance", h.getBalance)
		}
	}
}
