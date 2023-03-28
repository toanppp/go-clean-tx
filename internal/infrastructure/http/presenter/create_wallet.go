package presenter

type CreateWallet struct {
	Balance int64 `json:"balance" binding:"required,min=0"`
}
