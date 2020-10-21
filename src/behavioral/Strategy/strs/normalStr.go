package strs

type NormalStr struct {
	Prices
}

func (str *NormalStr) SetData(total int, price float64) {
	str.total = total
	str.price = price
}

func (str *NormalStr) Algorithm() float64 {
	return str.price * float64(str.total)
}
