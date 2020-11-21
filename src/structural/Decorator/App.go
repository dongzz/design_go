package main

import "fmt"

type Person interface {
	Show()
}

type Finery struct {
	Person
	name string
}

func (finery *Finery) Show() {
	if finery.Person != nil {
		finery.Person.Show()
	}
	fmt.Printf("%s 照镜子\n", finery.name)
}

type TShirt struct {
	Person
}

func (finery *TShirt) Show() {
	fmt.Printf("穿上了T恤\n")
	if finery.Person != nil {
		finery.Person.Show()
	}
}

type BigTrouser struct {
	Person
}

func (finery *BigTrouser) Show() {
	fmt.Printf("穿上了垮裤\n")
	if finery.Person != nil {
		finery.Person.Show()
	}
}

func main() {
	person := new(Person)
	fmt.Println("第一次")
	finery := &Finery{name: "dongz", Person: *person}
	finery.Show()
	fmt.Println("第二次")
	shirt := &TShirt{Person(finery)}
	shirt.Show()
	fmt.Println("第三次")
	trouser := BigTrouser{shirt}
	trouser.Show()
}
