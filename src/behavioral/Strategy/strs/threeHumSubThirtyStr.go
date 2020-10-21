package strs

type ThreeHumdrumSubThirtyStr struct {
	Prices
}

func (str *ThreeHumdrumSubThirtyStr) SetData(prices float64) {
	str.prices = prices
}

func (str *ThreeHumdrumSubThirtyStr) Algorithm() float64 {
	return str.prices - float64((int(str.prices/300))*30)
}
