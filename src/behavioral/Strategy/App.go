package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type CashContext struct {
	cash CashSuper
}

func (context *CashContext) SetCash(strategy int) {
	switch strategy {
	case 0:
		context.cash = &NormalCash{}
	case 1:
		context.cash = &EightCash{}
	case 2:
		context.cash = &SevenCash{}
	case 3:
		context.cash = &FiveCash{}
	case 4:
		context.cash = &FullSubCash{}
	}
}

func (context *CashContext) GetResult(money float64) float64 {
	return context.cash.AcceptCash(money)
}

type CashSuper interface {
	AcceptCash(money float64) float64
}

type NormalCash struct{}

func (cash *NormalCash) AcceptCash(money float64) float64 {
	return money
}

type EightCash struct{}

func (cash *EightCash) AcceptCash(money float64) float64 {
	return money * 0.8
}

type SevenCash struct{}

func (cash *SevenCash) AcceptCash(money float64) float64 {
	return money * 0.7
}

type FiveCash struct{}

func (cash *FiveCash) AcceptCash(money float64) float64 {
	return money * 0.5
}

type FullSubCash struct{}

func (cash *FullSubCash) AcceptCash(money float64) float64 {
	return money - float64(int(money)/300)*30
}

func main() {
	var total int
	var price float64
	var strategy int
	var err error
	strMap := make(map[int]string)
	strMap[0] = "正常收费"
	strMap[1] = "打八折"
	strMap[2] = "打七折"
	strMap[3] = "打五折"
	strMap[4] = "满300减30"

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入商品数量：")
	for scanner.Scan() {
		total, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("请输入正确的商品数量：")
		} else if total <= 0 {
			fmt.Println("商品数量需大于0：")
		} else {
			break
		}
	}

	fmt.Println("请输入商品价格：")
	for scanner.Scan() {
		price, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("请输入正确的商品价格：")
		} else if price <= 0 {
			fmt.Println("商品价格需大于0")
		} else {
			break
		}
	}

	fmt.Println("请选择折扣类型：")
	for k, v := range strMap {
		fmt.Printf("%d ---> %s \n", k, v)
	}
	fmt.Println("请选择：")
	for scanner.Scan() {
		strategy, err = strconv.Atoi(scanner.Text())
		if err != nil || strMap[strategy] == "" {
			fmt.Println("请输入正确的折扣类型：")
		} else {
			break
		}
	}

	context := &CashContext{}
	context.SetCash(strategy)

	fmt.Printf("单价：%f 数量：%d 折扣方式：%s 总价： %f\n", price, total, strMap[strategy], context.GetResult(price*float64(total)))
}
