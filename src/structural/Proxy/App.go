package main

import "fmt"

type GiveGift interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}

type Pursuit struct {
	GiveGift
	girlName string
}

func (p *Pursuit) GiveDolls() {
	fmt.Printf("%s, give you dolls\n", p.girlName)
}

func (p *Pursuit) GiveFlowers() {
	fmt.Printf("%s, give you flowers\n", p.girlName)
}

func (p *Pursuit) GiveChocolate() {
	fmt.Printf("%s, give you chocolate\n", p.girlName)
}

type PursuitProxy struct {
	Pursuit
	GiveGift
}

func (proxy *PursuitProxy) GiveDolls() {
	proxy.Pursuit.GiveDolls()
}

func (proxy *PursuitProxy) GiveFlowers() {
	proxy.Pursuit.GiveFlowers()
}

func (proxy *PursuitProxy) GiveChocolate() {
	proxy.Pursuit.GiveChocolate()
}

func main() {
	proxy := PursuitProxy{Pursuit: Pursuit{girlName: "caicai"}}

	proxy.GiveDolls()
	proxy.GiveFlowers()
	proxy.GiveChocolate()
}
