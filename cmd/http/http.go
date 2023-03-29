package main

import (
	"log"

	"github.com/gin-gonic/gin"
	router "github.com/toanppp/go-clean-tx/internal/infrastructure/http/gin"
	"github.com/toanppp/go-clean-tx/internal/infrastructure/repo/database"
	"github.com/toanppp/go-clean-tx/internal/usecase"
)

func main() {
	walletRepo := database.NewWalletRepo(database.NewDB())
	walletUseCase := usecase.NewWalletUseCase(walletRepo)

	engine := gin.New()

	router.RegisterHandle(engine, walletUseCase)

	if err := engine.Run(); err != nil {
		log.Print("gin.Engine.Run: ", err)
	}
}
