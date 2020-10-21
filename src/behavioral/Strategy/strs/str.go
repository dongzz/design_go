package strs

type Prices struct {
	total int
	price float64
}

type Str interface {
	SetData(total int, price float64)
	Algorithm() float64
}
