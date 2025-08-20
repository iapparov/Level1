package main

import "fmt"

func WhatType(v interface{}){ //interface{} == any
    switch v.(type){
        case int:
            fmt.Println("int")
        case string:
            fmt.Println("string")
        case bool:
            fmt.Println("bool")
        case chan int:
            fmt.Println("chan")
        default:
            fmt.Println("Неизвестная переменная")
    }
}

func main() {
    v1 := 10
    v2 := "WB"
    v3 := true
    v4 := make(chan int)

    WhatType(v1)
    WhatType(v2)
    WhatType(v3)
    WhatType(v4)
}

