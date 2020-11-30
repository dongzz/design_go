package main

import "fmt"

type Context struct {
	state State
}

func (context *Context) Request() {
	context.state.Handler(context)
	fmt.Println("当前状态：", context.state)
}

type State interface {
	Handler(context *Context)
}

type ConcreteStateA struct {
	name string
}

func (state *ConcreteStateA) Handler(context *Context) {
	context.state = &ConcreteStateB{name: "dongz"}
}

type ConcreteStateB struct {
	name string
}

func (state *ConcreteStateB) Handler(context *Context) {
	context.state = &ConcreteStateA{name: "caicai"}
}

func main() {
	context := Context{state: &ConcreteStateA{}}

	context.Request()
	context.Request()
	context.Request()
	context.Request()
	context.Request()
}
