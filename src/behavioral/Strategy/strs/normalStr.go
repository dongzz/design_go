package strs

type NormalStr struct {
	Prices
}

func (str *NormalStr) SetData(prices float64) {
	str.prices = prices
}

func (str *NormalStr) Algorithm() float64 {
	return str.prices
}
