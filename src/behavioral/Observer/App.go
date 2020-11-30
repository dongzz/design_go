package main

import (
	"fmt"
)

type Publisher interface {
	Attach(observer Observer)
	Detach(observer Observer)
	GetState() string
	Notify()
}

type Boss struct {
	observers []Observer
	state     string
}

func (publisher *Boss) Attach(observer Observer) {
	publisher.observers = append(publisher.observers, observer)
}

func (publisher *Boss) Detach(observer Observer) {
	target := publisher.observers[:0]
	for _, item := range publisher.observers {
		if item != observer {
			target = append(target, item)
		}
	}
	publisher.observers = target
}

func (publisher *Boss) GetState() string {
	return publisher.state
}

func (publisher *Boss) Notify() {
	for _, item := range publisher.observers {
		item.Update()
	}
}

type Observer interface {
	Update()
}

type ConcreteObserverA struct {
	publisher Publisher
}

func (observer *ConcreteObserverA) Update() {
	fmt.Printf("%s ,dongz 开始工作\n", observer.publisher.GetState())
}

type ConcreteObserverB struct {
	publisher Publisher
}

func (observer *ConcreteObserverB) Update() {
	fmt.Printf("%s ,pangpang 开始工作\n", observer.publisher.GetState())
}

func main() {
	boss := new(Boss)

	a := ConcreteObserverA{publisher: boss}
	boss.Attach(&a)

	b := ConcreteObserverB{publisher: boss}
	boss.Attach(&b)
	boss.state = "老板来啦！"

	boss.Notify()
}
