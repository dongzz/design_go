package comps

import "fmt"

// 复合构件
type Composite struct {
	Name        string
	Description string
	Arr         []Menu
}

func (c *Composite) GetName() string {
	return c.Name
}

func (c *Composite) GetPrice() float64 {
	panic("It is not an item.")
}

func (c *Composite) GetDescription() string {
	return c.Description
}

func (c *Composite) IsVegetarian() bool {
	panic("implement me")
}

func (c *Composite) CreateIterator() {
	panic("implement me")
}

func (c *Composite) Add(menu Menu) {
	c.Arr = append(c.Arr, menu)
}

func (c *Composite) Remove(menu Menu) {
	for i, v := range c.Arr {
		if v == menu {
			// 删除第i个元素,因为interface类型在golang中
			// 以地址的方式传递，所以可以直接比较进行删除
			// golang中只要记得byte,int,bool,string，数组，结构体，默认传值，其他的默认传地址即可
			c.Arr = append(c.Arr[:i], c.Arr[i+1:]...)
		}
	}
}

func (c *Composite) Display(depth int) {
	// 输出树形结构
	for i := 0; i < depth; i++ {
		fmt.Print("-")
	}
	fmt.Println(">", c.GetName())
	// 递归显示
	for _, com := range c.Arr {
		com.Display(depth + 1)
	}
}
