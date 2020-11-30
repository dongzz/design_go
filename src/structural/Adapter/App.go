package main

import "fmt"

type Player interface {
	Attack()
	Defense()
}

type Forwards struct {
	name string
}

func (player *Forwards) Attack() {
	fmt.Printf("前锋%s进攻\n", player.name)
}

func (player *Forwards) Defense() {
	fmt.Printf("前锋%s防守\n", player.name)
}

type Center struct {
	name string
}

func (player *Center) Attack() {
	fmt.Printf("中锋%s进攻\n", player.name)
}

func (player *Center) Defense() {
	fmt.Printf("中锋%s防守\n", player.name)
}

type Guards struct {
	name string
}

func (player *Guards) Attack() {
	fmt.Printf("后卫%s进攻\n", player.name)
}

func (player *Guards) Defense() {
	fmt.Printf("后卫%s防守\n", player.name)
}

type Translator struct {
	foreign ForeignCenter
}

func (translator *Translator) Attack() {
	translator.foreign.进攻()
}

func (translator *Translator) Defense() {
	translator.foreign.防守()
}

type ForeignCenter struct {
	name string
}

func (foreign *ForeignCenter) 进攻() {
	fmt.Printf("外籍中锋%s进攻\n", foreign.name)
}

func (foreign *ForeignCenter) 防守() {
	fmt.Printf("外籍中锋%s防守\n", foreign.name)
}

func main() {
	forwards := Forwards{name: "dongz"}
	forwards.Attack()

	guards := Guards{name: "caicai"}
	guards.Defense()

	translator := Translator{foreign: ForeignCenter{name: "姚明"}}
	translator.Attack()
	translator.Defense()
}
