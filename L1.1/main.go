package main

import "fmt"


type Human struct{
	Name string
	Age int
}

type Action struct{
	*Human // Можно и без указателя, но будет создаваться копия экземпляря Human
}

func main(){
	H := NewHuman("Ivan", 25)
	A := NewAction(H)
	A.Aging(100) //Увеличиваем возраст на 100
	fmt.Println(A.Age)
}

func NewHuman(name string, age int) *Human{
	return &Human{
		Name: name,
		Age: age,
	}
} //Инициализация экземпляра структуры Human

func NewAction(human *Human) *Action{
	return &Action{
		Human: human,
	}
} //Инициализация экземпляря структуры Action

func (h *Human) Aging(HowOld int){
	h.Age+=HowOld
}

/* 

Можно сделать структуру без указателя и инициализировать экземпляры без указателей. 

type Action struct{
	Human
}


Но в этом случае при 
A.Aging(100) //Увеличиваем возраст на 100

A.Age -> 125
H.Age -> 25

т.к. создастся копия human

*/