package strs

type Prices struct {
	prices float64
}

type Str interface {
	SetData(prices float64)
	Algorithm() float64
}
