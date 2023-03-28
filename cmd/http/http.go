package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/domain"
	router "github.com/toanppp/go-clean-tx/internal/infrastructure/http/gin"
	"github.com/toanppp/go-clean-tx/internal/infrastructure/repository/memory"
	"github.com/toanppp/go-clean-tx/internal/usecase"
)

func main() {
	walletRepository := memory.NewWalletRepository(map[int64]domain.Wallet{}, 0)
	walletUseCase := usecase.NewWalletUseCase(walletRepository)

	engine := gin.New()

	router.RegisterHandle(engine, walletUseCase)

	if err := engine.Run(); err != nil {
		log.Print("gin.Engine.Run: ", err)
	}
}
