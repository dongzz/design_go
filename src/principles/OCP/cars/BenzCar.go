package cars

type BenzCar struct {
	Name  string
	Price int
}

func (b BenzCar) GetName() string {
	return b.Name
}

func (b BenzCar) GetPrice() int {
	return b.Price
}
