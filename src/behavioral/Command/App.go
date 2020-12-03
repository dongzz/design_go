package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

type Barbecuer struct {
	muttonCount  int
	chickenCount int
}

func (barbecuer *Barbecuer) BakeMutton() {
	lock.Lock()
	if barbecuer.muttonCount > 0 {
		barbecuer.muttonCount--
		fmt.Printf("烤了一串肉串，还剩%d串肉串\n", barbecuer.muttonCount)
	} else {
		fmt.Println("肉串没有了")
	}
	lock.Unlock()
}

func (barbecuer *Barbecuer) BakeChicken() {
	lock.Lock()
	if barbecuer.chickenCount > 0 {
		barbecuer.chickenCount--
		fmt.Printf("烤了一串鸡翅，还剩%d串鸡翅\n", barbecuer.chickenCount)
	} else {
		fmt.Println("鸡翅没有了")
	}
	lock.Unlock()
}

type Command interface {
	Execute()
}

type BakeMuttonCommand struct {
	receiver *Barbecuer
}

func (command *BakeMuttonCommand) Execute() {
	command.receiver.BakeMutton()
}

type BakeChickenCommand struct {
	receiver *Barbecuer
}

func (command *BakeChickenCommand) Execute() {
	command.receiver.BakeChicken()
}

type Waiter struct {
	list []Command
}

func (waiter *Waiter) SetCommand(command Command) {
	waiter.list = append(waiter.list, command)
}

func (waiter *Waiter) RemoveCommand(command Command) {
	for i, c := range waiter.list {
		if c == command {
			waiter.list = append(waiter.list[:i], waiter.list[i+1:]...)
		}
	}
}

func (waiter *Waiter) Notify() {
	for _, command := range waiter.list {
		command.Execute()
	}
}

func main() {
	barbecuer := &Barbecuer{muttonCount: 10, chickenCount: 5}

	waiter := Waiter{}
	command := BakeMuttonCommand{receiver: barbecuer}
	waiter.SetCommand(&command)
	waiter.SetCommand(&BakeMuttonCommand{receiver: barbecuer})
	waiter.SetCommand(&BakeMuttonCommand{receiver: barbecuer})

	waiter.SetCommand(&BakeChickenCommand{receiver: barbecuer})
	waiter.SetCommand(&BakeChickenCommand{receiver: barbecuer})
	waiter.SetCommand(&BakeChickenCommand{receiver: barbecuer})
	waiter.SetCommand(&BakeChickenCommand{receiver: barbecuer})
	waiter.SetCommand(&BakeChickenCommand{receiver: barbecuer})
	waiter.SetCommand(&BakeChickenCommand{receiver: barbecuer})

	waiter.RemoveCommand(&command)
	waiter.Notify()
}
