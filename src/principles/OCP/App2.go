package main

import "fmt"
import "principles/OCP/cars"

func main() {
	var (
		list []cars.ICar
	)
	list = []cars.ICar{}
	list = append(list, &cars.FinanceBenzCar{BenzCar: cars.BenzCar{Name: "迈巴赫", Price: 99}})
	list = append(list, &cars.FinanceBenzCar{BenzCar: cars.BenzCar{Name: "AMG", Price: 200}})
	list = append(list, &cars.FinanceBenzCar{BenzCar: cars.BenzCar{Name: "V", Price: 40}})
	for _, v := range list {
		fmt.Println("车名:", v.GetName(), "\t价格:", v.GetPrice())
	}
}
