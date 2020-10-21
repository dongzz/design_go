package comps

type Menu interface {
	// 采购设备或者添加子部门
	Add(menu Menu)
	Remove(menu Menu)
	// 查询该节点下所有设备和部门
	GetName() string
	GetPrice() float64
	GetDescription() string
	IsVegetarian() bool
	CreateIterator()
	Display(depth int)
}
