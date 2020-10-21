package strs

type EightStr struct {
	Prices
}

func (str *EightStr) SetData(prices float64) {
	str.prices = prices
}

func (str *EightStr) Algorithm() float64 {
	return str.prices * 0.8
}
