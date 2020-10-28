package cars

type ICar interface {
	// 车名
	GetName() string
	// 价格
	GetPrice() int
}
