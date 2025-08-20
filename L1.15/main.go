package main

import(
  "strings"
)

var justString string

// createHugeString создаёт строку длиной size символов
func createHugeString(size int) string {
	return strings.Repeat("a", size) // Используем strings.Repeat для эффективного создания строки
}

func someFunc() {
  v := createHugeString(1 << 10) //2^10 == 10000000000
  justString = string([]byte(v[:100])) //В этом случае в justString мы скопируем первые 100 элментов от v
}

func main() {
  someFunc()
}


/* 
&lt;&lt; - вряд ли это "ошибка", но здесь должно быть <<

justString = v[:100] //Строки в го иммутабельны, логически мы должно взять первые 100 элементов из строки, 
но justString будет хранить ссылку на всю v, что может привести к утечкам памяти
*/