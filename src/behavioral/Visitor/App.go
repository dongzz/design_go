package main

import "fmt"

type State interface {
	GetManConclusion(man *Man)
	GetWomanConclusion(woman *Woman)
}

type Person interface {
	Accept(state State)
}

type Man struct {
	name string
}

func (person *Man) Accept(state State) {
	state.GetManConclusion(person)
}

type Woman struct {
	name string
}

func (person *Woman) Accept(state State) {
	state.GetWomanConclusion(person)
}

type Success struct {
	state string
}

func (state *Success) GetManConclusion(man *Man) {
	fmt.Printf("%s %s 时，背后多半有个默默付出的女人\n", man.name, state.state)
}

func (state *Success) GetWomanConclusion(woman *Woman) {
	fmt.Printf("%s %s 时，背后多半有一群默默付出的男人\n", woman.name, state.state)
}

type Failing struct {
	state string
}

func (state *Failing) GetManConclusion(man *Man) {
	fmt.Printf("%s %s 时，闷头喝酒，谁也不用劝\n", man.name, state.state)
}

func (state *Failing) GetWomanConclusion(woman *Woman) {
	fmt.Printf("%s %s 时，眼泪汪汪，谁也劝不动\n", woman.name, state.state)
}

type Loving struct {
	state string
}

func (state *Loving) GetManConclusion(man *Man) {
	fmt.Printf("%s %s 时，心花路放\n", man.name, state.state)
}

func (state *Loving) GetWomanConclusion(woman *Woman) {
	fmt.Printf("%s %s 时，小鹿乱撞\n", woman.name, state.state)
}

type ObjectStructure struct {
	list []Person
}

func (structure *ObjectStructure) Add(person Person) {
	structure.list = append(structure.list, person)
}

func (structure *ObjectStructure) Remove(person Person) {
	for i, p := range structure.list {
		if p == person {
			structure.list = append(structure.list[i:], structure.list[:i+1]...)
		}
	}
}

func (structure *ObjectStructure) Display(state State) {
	for _, person := range structure.list {
		person.Accept(state)
	}
}

func main() {
	structure := ObjectStructure{}

	man := Man{name: "男人"}
	woman := Woman{name: "女人"}

	structure.Add(&man)
	structure.Add(&woman)

	structure.Display(&Success{state: "成功"})
	structure.Display(&Failing{state: "失败"})
	structure.Display(&Loving{state: "恋爱"})
}
