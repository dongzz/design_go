package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Context struct {
	text string
}

type Expression interface {
	Interpreter(context *Context)
	Execute(key string, value float64)
}

func Interpreter(context *Context) (key string, value float64) {
	context.text = strings.TrimLeft(context.text, " ")
	if len(context.text) == 0 {
		return
	}
	key = context.text[0:1]
	context.text = context.text[2:]
	var err error
	for i, s := range context.text {
		if s == 32 {
			s2 := context.text[0:i]
			value, err = strconv.ParseFloat(s2, 64)
			if err != nil {
				return
			}
			context.text = context.text[i+1:]
			break
		}
	}
	return key, value
}

type Node struct {
}

func (expression *Node) Interpreter(context *Context) {
	expression.Execute(Interpreter(context))
}

func (expression *Node) Execute(key string, value float64) {
	var node = ""
	switch key {
	case "C":
		node = "1"
	case "D":
		node = "2"
	case "E":
		node = "3"
	case "F":
		node = "4"
	case "G":
		node = "5"
	case "A":
		node = "6"
	case "B":
		node = "7"
	}
	fmt.Printf("%s ", node)
}

type Scale struct {
}

func (expression *Scale) Interpreter(context *Context) {
	expression.Execute(Interpreter(context))
}

func (expression *Scale) Execute(key string, value float64) {
	var scale = ""
	switch value {
	case 1:
		scale = "低音"
	case 2:
		scale = "中音"
	case 3:
		scale = "高音"
	}
	fmt.Printf("%s ", scale)
}

func main() {
	context := Context{text: "O 2 E 0.5 G 0.5 A 3 E 0.5 G 0.5 D 3 E 0.5 G 0.5 A 0.5 O 3 C 1 O 2 A 0.5 G 1 C 0.5 E 0.5 D 3 "}

	var expression Expression

	for len(context.text) > 0 {
		str := context.text[0:1]
		switch str {
		case "O":
			expression = &Scale{}
		case "C":
			fallthrough
		case "D":
			fallthrough
		case "E":
			fallthrough
		case "F":
			fallthrough
		case "G":
			fallthrough
		case "A":
			fallthrough
		case "B":
			fallthrough
		case "P":
			expression = &Node{}
		default:
			expression = nil
		}
		if expression != nil {
			expression.Interpreter(&context)
		}
	}
}
