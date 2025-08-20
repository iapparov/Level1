package main

import "fmt"

func main() {
    words := []string{"cat", "cat", "dog", "cat", "tree"}

    set := make(map[string]bool) // Создаём множество через map
    for _, word := range words {
        set[word] = true
    }

    for word := range set { // Вывод элементов. Т.к. это мапа – порядок вывода может отличаться
        fmt.Printf("%s ", word)
    }
}