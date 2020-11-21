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

type OperatorSimpleFactory struct{}

func (factory *OperatorSimpleFactory) CreateOperator(operatorType string) Operator {
	switch operatorType {
	case "+":
		return &AddOperator{}
	case "-":
		return &SubOperator{}
	default:
		return nil
	}
}

func main() {
	var numberA, numberB float64
	var operatorType string
	var err error
	var operator Operator
	factory := &OperatorSimpleFactory{}
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
			operator = factory.CreateOperator(operatorType)
			if operator == nil {
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
	operator.SetData(numberA, numberB)

	fmt.Printf("%f %s %f = %f", numberA, operatorType, numberB, operator.GetResult())
}
