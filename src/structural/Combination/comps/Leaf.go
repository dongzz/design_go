package comps

import "fmt"

type Leaf struct {
	Name        string
	Vegetarian  bool
	Description string
	Price       float64
}

func (l *Leaf) Add(menu Menu) {
	panic("Leaf Node can not add")
}

func (l *Leaf) Remove(menu Menu) {
	panic("Leaf Node can not remove")
}

func (l *Leaf) GetName() string {
	return l.Name
}

func (l *Leaf) GetPrice() float64 {
	return l.Price
}

func (l *Leaf) IsVegetarian() bool {
	return l.Vegetarian
}

func (l *Leaf) GetDescription() string {
	return l.Description
}

func (l *Leaf) CreateIterator() {
	panic("implement me")
}

func (l *Leaf) Display(depth int) {
	// 输出树形结构的叶子结点，这里直接输出设备名

	fmt.Print("|")
	for i := 0; i < depth; i++ {
		fmt.Print("-")
	}
	fmt.Println(">", l.GetName())
}

func (l *Leaf) SetName(name string) {
	l.Name = name
}
