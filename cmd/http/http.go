package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/domain"
	"github.com/toanppp/go-clean-tx/internal/wallet/infrastructure/http/gin"
	"github.com/toanppp/go-clean-tx/internal/wallet/infrastructure/repository/in_memory"
	"github.com/toanppp/go-clean-tx/internal/wallet/use_case"
)

func main() {
	walletRepository := wallet_repository_in_memory.NewWalletRepository(map[int64]domain.Wallet{}, 0)
	walletUseCase := wallet_use_case.NewWalletUseCase(walletRepository)

	engine := gin.New()

	wallet := engine.Group("wallet")
	wallet_http_gin.RegisterWalletGinHTTP(wallet, walletUseCase)

	if err := engine.Run(); err != nil {
		log.Print("gin.Engine.Run: ", err)
	}
}
