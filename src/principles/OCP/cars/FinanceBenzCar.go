package cars

type FinanceBenzCar struct {
	BenzCar
}

func (b FinanceBenzCar) GetPrice() int {
	// 获取原价
	selfPrice := b.Price
	var finance int
	if selfPrice >= 100 {
		finance = selfPrice + selfPrice*5/100
	} else if selfPrice >= 50 {
		finance = selfPrice + selfPrice*2/100
	} else {
		finance = selfPrice
	}
	return finance
}
