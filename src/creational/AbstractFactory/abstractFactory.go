package main

import "fmt"

type Factory interface {
	newTV() TV
	newRefrigerator() Refrigerator
}

type TV interface {
	TVShow()
}

type Refrigerator interface {
	RefrigeratorShow()
}

//////////////////////

type TCLFactory struct{}

func (factory *TCLFactory) newTV() TV {
	return &TCLTV{name: "TCL TV200", price: 20000}
}

func (factory *TCLFactory) newRefrigerator() Refrigerator {
	return &TCLRefrigerator{name: "TCL Refrigerator200", price: 2000}
}

type MeiDFactory struct{}

func (factory *MeiDFactory) newTV() TV {
	return &MeiDTV{name: "MeiD TV 3300", price: 100000}
}

func (factory *MeiDFactory) newRefrigerator() Refrigerator {
	return &MeiDRefrigerator{name: "MeiD Refrigerator 1200", price: 11000}
}

////////////////////////

type TCLTV struct {
	name  string
	price float64
}

func (tv *TCLTV) TVShow() {
	fmt.Printf("TCL TV name is %s, price is %f\n", tv.name, tv.price)
}

type MeiDTV struct {
	name  string
	price float64
}

func (tv *MeiDTV) TVShow() {
	fmt.Printf("MeiD TV name is %s, price is %f\n", tv.name, tv.price)
}

///////////////////

type TCLRefrigerator struct {
	name  string
	price float64
}

type MeiDRefrigerator struct {
	name  string
	price float64
}

func (refigerator *TCLRefrigerator) RefrigeratorShow() {
	fmt.Printf("TCL Refrigerator name is %s, price is %f\n", refigerator.name, refigerator.price)
}

func (refigerator *MeiDRefrigerator) RefrigeratorShow() {
	fmt.Printf("MeiD Refrigerator name is %s, price is %f\n", refigerator.name, refigerator.price)
}

func main() {
	f := new(MeiDFactory)
	tv := f.newTV()
	tv.TVShow()
	ref := f.newRefrigerator()
	ref.RefrigeratorShow()

	f2 := new(TCLFactory)
	tv = f2.newTV()
	tv.TVShow()
	ref = f2.newRefrigerator()
	ref.RefrigeratorShow()
}
