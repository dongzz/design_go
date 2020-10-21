package main

import "structural/Combination/comps"

func main() {
	root := comps.Composite{
		Name: "和平饭店",
		Arr:  make([]comps.Menu, 0),
	}

	branchLevel21 := comps.Composite{
		Name: "招牌菜",
		Arr:  make([]comps.Menu, 0),
	}
	branchLevel21.Add(&comps.Leaf{Name: "红烧肉", Description: "精五花", Price: 20.0})
	branchLevel21.Add(&comps.Leaf{Name: "醋溜土豆丝", Vegetarian: true, Description: "新鲜", Price: 10.0})
	branchLevel21.Add(&comps.Leaf{Name: "京酱肉丝", Description: "鲜肉", Price: 30.0})

	root.Add(&branchLevel21)

	// 并列的二级节点
	branchLevel22 := comps.Composite{
		Name: "家常菜",
		Arr:  make([]comps.Menu, 0),
	}
	branchLevel22.Add(&comps.Leaf{Name: "辣椒炒肉", Description: "1", Price: 14.2})
	branchLevel22.Add(&comps.Leaf{Name: "杂拌", Description: "1234", Price: 15})
	branchLevel22.Add(&comps.Leaf{Name: "回锅肉", Description: "2134", Price: 30})

	branchLevel221 := comps.Composite{
		Name: "麻辣烫",
		Arr:  make([]comps.Menu, 0),
	}
	branchLevel221.Add(&comps.Leaf{Name: "豆蔻", Vegetarian: true, Description: "1", Price: 0.5})
	branchLevel221.Add(&comps.Leaf{Name: "腐竹", Vegetarian: true, Description: "1", Price: 0.5})
	branchLevel22.Add(&branchLevel221)

	root.Add(&branchLevel22)

	root.Display(1)

	root.Remove(&branchLevel22)
	root.Display(1)
}
