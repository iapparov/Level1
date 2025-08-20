package main

import (
	"fmt"
)


func main(){
	fmt.Println("Количество элеменов в 1 слайсе:")
	var n int
	fmt.Scan(&n)
	arr1 := make([]int,n)
	fmt.Println("Введите элементы для 1 слайса:")
	for i := 0; i<n; i++{
		fmt.Scan(&arr1[i])
	}
	fmt.Println("Количество элеменов во 2 слайсе:")
	fmt.Scan(&n)
	arr2 := make([]int,n)
	fmt.Println("Введите элементы для 2 слайса:")
	for i := 0; i<n; i++{
		fmt.Scan(&arr2[i])
	}
	//arr1 := []int{1,2,3}
	//arr2 := []int{2,3,4}
	ans := make([]int,0)
	is := make(map[int]bool)
	for _, elem := range arr1{
		is[elem]=true
	}
	for _, elem :=range arr2{
		if is[elem]{
			ans = append(ans, elem)
			is[elem] = false //Для того, чтобы не добавлять повторяющиеся пересечения в слайс
		}
	}
	for _, elem := range ans{
		fmt.Printf("%d ", elem)
	}
	fmt.Printf("\n")
}