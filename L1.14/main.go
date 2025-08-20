package main

import (
    "fmt"
    "reflect")

func WhatType(v interface{}){ //interface{} == any
    switch v.(type){
        case int:
            fmt.Println("int")
        case string:
            fmt.Println("string")
        case bool:
            fmt.Println("bool")
        default:    
            rv := reflect.ValueOf(v)// проверяем через reflect, если это канал любого типа
            if rv.Kind() == reflect.Chan {
                fmt.Println("chan")
            } else {
                fmt.Println("Неизвестная переменная")
            }
    }
}

func main() {
    v1 := 10
    v2 := "WB"
    v3 := true
    v4 := make(chan int)
    v5 := make(chan string)

    WhatType(v1)
    WhatType(v2)
    WhatType(v3)
    WhatType(v4)
    WhatType(v5)
}
/*

Вообще исходя из такси нам намекают на использование switch v.(type)

но здесь как будто бы вообще можно воспользоваться 
rv := reflect.Valueof(v)
fmt.Println(rv.Kind())

*/
