package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototypeA struct {
	name string
	age  int
	sex  string
}

func (prototype *ConcretePrototypeA) Clone() Prototype {
	return &ConcretePrototypeA{name: prototype.name, age: prototype.age, sex: prototype.sex}
}

type ConcretePrototypeB struct {
	ConcretePrototypeA
	weight float64
}

func (prototype *ConcretePrototypeB) Clone() Prototype {
	return &ConcretePrototypeB{weight: prototype.weight, ConcretePrototypeA: *prototype.ConcretePrototypeA.Clone().(*ConcretePrototypeA)}
}

func main() {
	prototypeA := ConcretePrototypeA{name: "dongz", age: 28, sex: "male"}
	prototypeA2 := prototypeA.Clone()
	fmt.Println(prototypeA)
	fmt.Println(prototypeA2)

	prototypeB := ConcretePrototypeB{weight: 62, ConcretePrototypeA: prototypeA}
	prototypeB2 := prototypeB.Clone().(*ConcretePrototypeB)
	prototypeB2.weight = 54
	prototypeB2.name = "caicai"
	prototypeB2.sex = "female"
	prototypeB2.age = 24

	fmt.Println(prototypeB)
	fmt.Println(prototypeB2)
}
