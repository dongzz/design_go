package main

import "fmt"

type InterfaceClass interface {
	begin()
	end()
	DoSomething(class InterfaceClass)
	DoWork()
}

type AbstractClass struct{}

func (class *AbstractClass) begin() {
	fmt.Println("begin")
}

func (class *AbstractClass) end() {
	fmt.Println("end")
}

func (class *AbstractClass) DoWork() {}

func (class *AbstractClass) DoSomething(do InterfaceClass) {
	class.begin()
	fmt.Println("doSomething...")
	do.DoWork()
	class.end()
}

type ConcreteClassA struct {
	AbstractClass
}

func (class *ConcreteClassA) DoWork() {
	fmt.Println("Do with Concrete option A")
}

func (class *ConcreteClassA) DoSomething(do InterfaceClass) {
	class.AbstractClass.DoSomething(class)
}

type ConcreteClassB struct {
	AbstractClass
}

func (class *ConcreteClassB) DoWork() {
	fmt.Println("Do with Concrete option B")
}

func (class *ConcreteClassB) DoSomething(do InterfaceClass) {
	class.AbstractClass.DoSomething(class)
}

func main() {
	var class InterfaceClass
	class = &ConcreteClassA{AbstractClass{}}
	class.DoSomething(class)

	class = &ConcreteClassB{AbstractClass{}}
	class.DoSomething(class)
}
