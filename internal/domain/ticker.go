package domain

type Ticker struct {
	Symbol string
	Price  float64
}

type TickerRepository interface {
	FetchPrice(symbol string) (string, error)
}
