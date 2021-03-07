package crypto

type currencyModel struct {
	BRL int `binding:"required"`
	EUR int `binding:"required"`
	CAD int `binding:"required"`
}
