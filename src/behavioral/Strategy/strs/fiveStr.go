package strs

type FiveStr struct {
	Prices
}

func (str *FiveStr) SetData(prices float64) {
	str.prices = prices
}

func (str *FiveStr) Algorithm() float64 {
	return str.prices * 0.5
}
