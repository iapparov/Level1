package main

import (
	"bufio"
	"fmt"
	"os"
)

func Reverse(s []rune, low int, high int){
    for low < high{
        s[low], s[high] = s[high], s[low]
        low++
        high--
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    word := scanner.Text()
    s := []rune(word)
    
    Reverse(s, 0, len(s)-1) // сначала реверсим всю строку в принципе
    var start int
    for idx := 0; idx < len(s); idx++{ //реверсим каждое слово
        if s[idx] == ' '{
            Reverse(s, start, idx-1)
            start = idx+1
        }
    }
    Reverse(s, start, len(s)-1) // реверсим последнее слово
    fmt.Print(string(s))
}

/*
Можно конечно было обойтись простым slices.Reverse(s) он аналогично переворачивает слайс рун
*/