package main

import "fmt"

type SubSystemOne struct{}

func (system SubSystemOne) SubMethodA() {
	fmt.Println("系统一方法")
}

type SubSystemTwo struct{}

func (system SubSystemTwo) SubMethodB() {
	fmt.Println("系统二方法")
}

type SubSystemThree struct{}

func (system SubSystemThree) SubMethodC() {
	fmt.Println("系统三方法")
}

type SubSystemFour struct{}

func (system SubSystemFour) SubMethodD() {
	fmt.Println("系统四方法")
}

type Facade struct {
	one   SubSystemOne
	two   SubSystemTwo
	three SubSystemThree
	four  SubSystemFour
}

func (facade *Facade) MethodA() {
	fmt.Println("方法 A")
	facade.one.SubMethodA()
	facade.three.SubMethodC()
	facade.four.SubMethodD()
}

func (facade *Facade) MethodB() {
	fmt.Println("方法 B")
	facade.two.SubMethodB()
}

func main() {
	facade := Facade{one: SubSystemOne{}, two: SubSystemTwo{}, three: SubSystemThree{}, four: SubSystemFour{}}

	facade.MethodA()
	facade.MethodB()
}
