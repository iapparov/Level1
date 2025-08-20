package main

import "fmt"

func main() {
    var numb1, numb2 int
    fmt.Scan(&numb1, &numb2)

    numb1, numb2 = numb2, numb1 //множественное присваивание
    fmt.Println(numb1, numb2)

    // Сложение/вычитание 
    numb1 = numb1 + numb2
    numb2 = numb1 - numb2
    numb1 = numb1 - numb2

    fmt.Println(numb1, numb2)

    //XOR обмен
    numb1 = numb1 ^ numb2
    numb2 = numb1 ^ numb2
    numb1 = numb1 ^ numb2
    fmt.Println(numb1, numb2)
}

