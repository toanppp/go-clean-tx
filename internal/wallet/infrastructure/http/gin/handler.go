package gin

import (
	"github.com/toanppp/go-clean-tx/internal/port"
)

type walletHandler struct {
	walletUseCase port.WalletUseCase
}
