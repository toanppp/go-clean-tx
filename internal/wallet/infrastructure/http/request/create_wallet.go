package wallet_http_request

type CreateWallet struct {
	Balance int64 `json:"balance" binding:"required,min=0"`
}
