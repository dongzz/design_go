package main

import "fmt"

type Iterator interface {
	First() interface{}
	Next() interface{}
	IsDone() bool
	CurrentItem() interface{}
}

type Aggregate interface {
	CreateIterator() Iterator
}

type ConcreteIterator struct {
	aggregate ConcreteAggregate
	current   int
}

func (iterator *ConcreteIterator) First() interface{} {
	return iterator.aggregate.list[0]
}

func (iterator *ConcreteIterator) CurrentItem() interface{} {
	return iterator.aggregate.list[iterator.current]
}

func (iterator *ConcreteIterator) Next() interface{} {
	iterator.current++
	if iterator.current < len(iterator.aggregate.list) {
		return iterator.aggregate.list[iterator.current]
	} else {
		return nil
	}
}

func (iterator *ConcreteIterator) IsDone() bool {
	return iterator.current >= len(iterator.aggregate.list)
}

type ConcreteAggregate struct {
	list []string
}

func (aggregate *ConcreteAggregate) CreateIterator() Iterator {
	return &ConcreteIterator{aggregate: *aggregate, current: 0}
}

func main() {
	strings := []string{"dongz", "caicai", "pangpang"}
	aggregate := ConcreteAggregate{list: strings}

	iterator := aggregate.CreateIterator()

	for !iterator.IsDone() {
		fmt.Println(iterator.CurrentItem())
		iterator.Next()
	}
}
