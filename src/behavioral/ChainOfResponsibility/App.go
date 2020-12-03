package main

import "fmt"

type Handler interface {
	SetSuccesser(successer Handler)
	HandlerRequest(state int)
}

type ConcreteHandlerA struct {
	successer Handler
}

func (handler *ConcreteHandlerA) SetSuccesser(successer Handler) {
	handler.successer = successer
}

func (handler *ConcreteHandlerA) HandlerRequest(state int) {
	if state < 5 {
		fmt.Println("ConcreteHandler A 处理")
	} else if handler.successer != nil {
		handler.successer.HandlerRequest(state)
	}
}

type ConcreteHandlerB struct {
	successer Handler
}

func (handler *ConcreteHandlerB) SetSuccesser(successer Handler) {
	handler.successer = successer
}

func (handler *ConcreteHandlerB) HandlerRequest(state int) {
	if state >= 5 {
		fmt.Println("ConcreteHandler B 处理")
	} else if handler.successer != nil {
		handler.successer.HandlerRequest(state)
	}
}

func main() {
	a := ConcreteHandlerA{nil}

	b := ConcreteHandlerB{successer: &a}

	num := []int{1, 5, 2, 6, 1, 2, 3, 88, 6, 8, 10}
	for _, i := range num {
		b.HandlerRequest(i)
	}
}
