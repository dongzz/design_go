package main

import "fmt"

type Mediator interface {
	Add(colleague Colleague)
	Remove(colleague Colleague)
	Send(message string, colleague Colleague)
}

type Colleague interface {
	Send(message string)
	Notify(message string)
}

type ConcreteMediator struct {
	colleagues []Colleague
}

func (mediator *ConcreteMediator) Add(colleague Colleague) {
	mediator.colleagues = append(mediator.colleagues, colleague)
}

func (mediator *ConcreteMediator) Remove(colleague Colleague) {
	for i, c := range mediator.colleagues {
		if c == colleague {
			mediator.colleagues = append(mediator.colleagues[:i], mediator.colleagues[i+1:]...)
		}
	}
}

func (mediator *ConcreteMediator) Send(message string, colleague Colleague) {
	for _, c := range mediator.colleagues {
		if c != colleague {
			c.Notify(message)
		}
	}
}

type ConcreteColleagueA struct {
	name     string
	mediator Mediator
}

func (colleague *ConcreteColleagueA) Send(message string) {
	colleague.mediator.Send(message, colleague)
}

func (colleague *ConcreteColleagueA) Notify(message string) {
	fmt.Printf("同事 A 收到消息 ： %s\n", message)
}

type ConcreteColleagueB struct {
	name     string
	mediator Mediator
}

func (colleague *ConcreteColleagueB) Send(message string) {
	colleague.mediator.Send(message, colleague)
}

func (colleague *ConcreteColleagueB) Notify(message string) {
	fmt.Printf("同事 B 收到消息 ： %s\n", message)
}

type ConcreteColleagueC struct {
	name     string
	mediator Mediator
}

func (colleague *ConcreteColleagueC) Send(message string) {
	colleague.mediator.Send(message, colleague)
}

func (colleague *ConcreteColleagueC) Notify(message string) {
	fmt.Printf("同事 C 收到消息 ： %s\n", message)
}

func main() {
	mediator := ConcreteMediator{}
	a := ConcreteColleagueA{mediator: &mediator}
	b := ConcreteColleagueB{mediator: &mediator}
	c := ConcreteColleagueC{mediator: &mediator}

	mediator.Add(&a)
	mediator.Add(&b)
	mediator.Add(&c)

	a.Send("hello")

	b.Send("nice to meet you")

	mediator.Remove(&c)

	a.Send("e....")
}
