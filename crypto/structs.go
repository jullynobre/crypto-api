package crypto

type currencyQuote struct {
	BRL int `binding:"required"`
	EUR int `binding:"required"`
	CAD int `binding:"required"`
}

type time struct {
	Updated    string `binding:"required"`
	UpdatedISO string `binding:"required"`
	Updateduk  string `binding:"required"`
}

type currencyRate struct {
	Code        string  `binding:"required"`
	Rate        string  `binding:"required"`
	Description string  `binding:"required"`
	RateFloat   float32 `json:"rate_float" binding:"required"`
}

type bpi struct {
	USD currencyRate `binding:"required"`
	BTC currencyRate `binding:"required"`
	BRL currencyRate
	EUR currencyRate
	CAD currencyRate
}

type cryptoResponse struct {
	Time       time   `binding:"required"`
	Disclaimer string `binding:"required"`
	BPI        bpi    `binding:"required"`
}