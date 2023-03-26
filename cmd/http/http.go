package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/domain"
	walletGin "github.com/toanppp/go-clean-tx/internal/wallet/infrastructure/http/gin"
	"github.com/toanppp/go-clean-tx/internal/wallet/infrastructure/repository/memory"
	"github.com/toanppp/go-clean-tx/internal/wallet/usecase"
)

func main() {
	walletRepository := memory.NewWalletRepository(map[int64]domain.Wallet{}, 0)
	walletUseCase := usecase.NewWalletUseCase(walletRepository)

	engine := gin.New()

	wallet := engine.Group("wallet")
	walletGin.RegisterWalletGinHTTP(wallet, walletUseCase)

	if err := engine.Run(); err != nil {
		log.Print("gin.Engine.Run: ", err)
	}
}
