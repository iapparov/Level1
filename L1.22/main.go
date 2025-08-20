package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Определяем два больших числа
	a := big.NewInt(734534957934794123) 
	b := big.NewInt(734534957934794123) 

	// Создаем переменные для результата
	sum := big.NewInt(0)
	min := big.NewInt(0)
	mul := big.NewInt(0)
	div := big.NewInt(0)

	sum.Add(a, b)
	min.Sub(a, b)
	mul.Mul(a, b)
	div.Div(a, b)

	fmt.Println("a + b", sum)
	fmt.Println("a - b", min)
	fmt.Println("a * b", mul) //здесь кстати хороший пример для использования bigint
	fmt.Println("a / b", div)
}