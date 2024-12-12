package main

import (
	"fmt"
	"github.com/shopspring/decimal" // Используем пакет decimal для работы с большими числами
)

func main() {

	value1, _ := decimal.NewFromString("4194304")
	value2, _ := decimal.NewFromString("4194310")

	// Используем встроенные методы для осуществления арифметических операций
	mul := value1.Mul(value2)
	div := value1.Div(value2)
	sub := value1.Sub(value2)
	sum := decimal.Sum(value1, value2)

	fmt.Println(value1)
	fmt.Println(value2)

	fmt.Println(mul)
	fmt.Println(div)
	fmt.Println(sub)
	fmt.Println(sum)

}
