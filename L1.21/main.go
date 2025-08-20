package main

import (
	"fmt"
)

type Mathematic interface{ //Новый интерфейс
    Plus(a int, b int) int
}

type Arithmetic struct{} //Старая структура

func (s *Arithmetic) Slozhenie(a int, b int) int{
    return a+b
}

type ArithmeticAdapter struct{
    *Arithmetic
}

func (s* ArithmeticAdapter) Plus(a int, b int) int{
    return s.Arithmetic.Slozhenie(a,b)
}

func main() {
    var m Mathematic
    m = &ArithmeticAdapter{Arithmetic: &Arithmetic{}}
    fmt.Println(m.Plus(1,2))
}

/*
+: Позволяет использовать старый код не переписывая его. 
-: Увеличивает количество оберток (больше кода) и можно легко запутаться

Как вариант можно написать адаптер для интеграции с внешним API который присылает данным в XML, а нам нужен JSON

*/