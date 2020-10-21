package strs

type SevenStr struct {
	Prices
}

func (str *SevenStr) SetData(prices float64) {
	str.prices = prices
}

func (str *SevenStr) Algorithm() float64 {
	return str.prices * 0.7
}
