package main

import "fmt"

type PersonBuilder interface {
	BuilderHead()
	BuilderBody()
	BuilderFoot()
}

type FatPersonBuilder struct{}

func (builder *FatPersonBuilder) BuilderHead() {
	fmt.Println("肥头大耳")
}

func (builder *FatPersonBuilder) BuilderBody() {
	fmt.Println("大肚便便")
}

func (builder *FatPersonBuilder) BuilderFoot() {
	fmt.Println("大象腿")
}

type ThinPersonBuilder struct{}

func (builder *ThinPersonBuilder) BuilderHead() {
	fmt.Println("尖嘴猴腮")
}

func (builder *ThinPersonBuilder) BuilderBody() {
	fmt.Println("瘦不拉几")
}

func (builder *ThinPersonBuilder) BuilderFoot() {
	fmt.Println("骨瘦如柴")
}

type PersonBuilderDirector struct {
	builder PersonBuilder
}

func (director *PersonBuilderDirector) CreatePerson() {
	director.builder.BuilderHead()
	director.builder.BuilderBody()
	director.builder.BuilderFoot()
}

func main() {
	var builder PersonBuilder
	builder = new(FatPersonBuilder)
	director := PersonBuilderDirector{builder: builder}
	director.CreatePerson()

	builder = new(ThinPersonBuilder)
	director = PersonBuilderDirector{builder: builder}
	director.CreatePerson()
}
