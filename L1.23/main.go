package main

import (
	"fmt"
)

func main() {
	s := []int{0,1,2,3,4,5,6,7,8}
	i:=5
	copy(s[i:], s[i+1:]) //удаляем элемент
	s[len(s)-1]=0 //Помогаем GC
	s = s[:len(s)-1] //уменьшаем длинну слайса на 1
	fmt.Println(s)

	//Еще можно так
	s1 := []int{0,1,2,3,4,5,6,7,8}
	s1 = append(s1[:i], s1[i+1:]...)
	fmt.Println(s1)
}