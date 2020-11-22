package main

import (
	"fmt"
	"os"
	"strconv"
)
import "bufio"

type Operator interface {
	GetResult() float64
	SetData(numberA, numberB float64)
}

type AddOperator struct {
	NumberA float64
	NumberB float64
}

func (a AddOperator) GetResult() float64 {
	return a.NumberA + a.NumberB
}

func (a *AddOperator) SetData(numberA, numberB float64) {
	a.NumberA = numberA
	a.NumberB = numberB
}

type SubOperator struct {
	NumberA float64
	NumberB float64
}

func (a SubOperator) GetResult() float64 {
	return a.NumberA - a.NumberB
}

func (a *SubOperator) SetData(numberA, numberB float64) {
	a.NumberA = numberA
	a.NumberB = numberB
}

type OperatorFactory interface {
	CreateOperator() Operator
}

type AddOperatorFactory struct{}

func (factory *AddOperatorFactory) CreateOperator() Operator {
	return &AddOperator{}
}

type SubOperatorFactory struct{}

func (factory *SubOperatorFactory) CreateOperator() Operator {
	return &SubOperator{}
}

func main() {
	var numberA, numberB float64
	var operatorType string
	var err error
	var factory OperatorFactory
	fmt.Println("请输入第一个数：")
	in := bufio.NewScanner(os.Stdin)
	for {
		if in.Scan() {
			numberA, err = strconv.ParseFloat(in.Text(), 64)
			if err != nil {
				fmt.Println("输入错误, 请重新输入！")
			} else {
				break
			}
		}
	}
	fmt.Println("请输入运算符：+、-")
	for {
		if in.Scan() {
			operatorType = in.Text()
			if operatorType == "+" {
				factory = &AddOperatorFactory{}
			} else if operatorType == "-" {
				factory = &SubOperatorFactory{}
			}
			if factory == nil {
				fmt.Println("请输入正确的运算符")
			} else {
				break
			}
		}
	}
	fmt.Println("请输入第二个数：")
	for {
		if in.Scan() {
			numberB, err = strconv.ParseFloat(in.Text(), 64)
			if err != nil {
				fmt.Println("输入错误, 请重新输入！")
			} else {
				break
			}
		}
	}
	operator := factory.CreateOperator()
	operator.SetData(numberA, numberB)

	fmt.Printf("%f %s %f = %f", numberA, operatorType, numberB, operator.GetResult())
}
