package main

import (
	"bufio"
	"creational/SimpleFactory/operates"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numOne, numTwo float64
	var strOperate string
	var err error
	var operate operates.Ops

	factory := new(operates.OpsFactory)

	fmt.Println("请输入第一个参数：")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		numOne, err = strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			fmt.Println("输入的格式错误，请重新输入：")
		} else {
			goto OPS
		}
	}

OPS:
	fmt.Println("请输入运算符（+，-，*，/）：")
	for scanner.Scan() {
		strOperate = scanner.Text()
		operate = factory.CreateOperate(strOperate)
		if operate == nil {
			fmt.Println("输入的格式错误，请重新输入 (+，-，*，/)：")
		} else {
			goto PARAMS
		}
	}

PARAMS:
	fmt.Println("请输入第二个参数：")
	for scanner.Scan() {
		numTwo, err = strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			fmt.Println("输入的格式错误，请重新输入：")
		} else {
			break
		}
	}

	operate.SetData(numOne, numTwo)
	result := operate.GetResult()
	fmt.Printf("%f %s %f = %f\n", numOne, strOperate, numTwo, result)
}
