package entity

type Payment struct {
	Transaction  string `json:"transaction" validate:"required"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency" validate:"iso4217"`
	Provider     string `json:"provider"`
	Amount       uint   `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost uint   `json:"delivery_cost"`
	Total        uint64 `json:"goods_total"`
	CustomFee    uint   `json:"custom_fee"`
}
