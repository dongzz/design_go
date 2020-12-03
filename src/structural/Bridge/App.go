package main

import "fmt"

type Implementor interface {
	OperationImpl()
}

type ConcreteImplA struct{}

func (impl *ConcreteImplA) OperationImpl() {
	fmt.Println("具体实现 A")
}

type ConcreteImplB struct{}

func (impl *ConcreteImplB) OperationImpl() {
	fmt.Println("具体实现 B")
}

type Abstraction struct {
	impl Implementor
}

func (abstract *Abstraction) Operation() {
	abstract.impl.OperationImpl()
}

func (abstract *Abstraction) SetImplementor(impl Implementor) {
	abstract.impl = impl
}

type RefinedAbstraction struct {
	Abstraction
}

func main() {
	abstraction := RefinedAbstraction{}
	abstraction.SetImplementor(&ConcreteImplA{})
	abstraction.Operation()
	abstraction.SetImplementor(&ConcreteImplB{})
	abstraction.Operation()
}
