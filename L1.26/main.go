package main

import (
	"fmt"
	"unicode"
)


func main() {
	var word string
	fmt.Scan(&word)
	unique := make(map[rune]bool)
	for _, elem := range word{
		elem = unicode.ToLower(elem) //Решаем проблему с регистром переводом все в строчные
		if unique[elem]{ // если хотя бы один элемент не уникальный, то можно вернуть false и завершить выполнение
			fmt.Println("false")
			return
		} else {
			unique[elem] = true
		}
	}
	fmt.Println("true")
}