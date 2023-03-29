package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/domain"
	router "github.com/toanppp/go-clean-tx/internal/infrastructure/http/gin"
	"github.com/toanppp/go-clean-tx/internal/infrastructure/repo/memory"
	"github.com/toanppp/go-clean-tx/internal/usecase"
)

func main() {
	walletRepo := memory.NewWalletRepo(map[int64]domain.Wallet{}, 0)
	walletUseCase := usecase.NewWalletUseCase(walletRepo)

	engine := gin.New()

	router.RegisterHandle(engine, walletUseCase)

	if err := engine.Run(); err != nil {
		log.Print("gin.Engine.Run: ", err)
	}
}
