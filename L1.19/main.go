package main

import (
	"fmt"
)



func main() {
    var word string
    fmt.Scan(&word)
    s := []rune(word)
    for idx := 0; idx< len(s)/2; idx++{
        s[idx], s[len(s)-1-idx] = s[len(s)-1-idx], s[idx]
    } // Это ручной вариант
    fmt.Println(string(s))
}

/*
Можно конечно было обойтись простым slices.Reverse(s) он аналогично переворачивает слайс рун
*/