package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func Worker(r chan int, id int, wg *sync.WaitGroup){
	defer wg.Done()
	for v := range r{ //Если канал закроется, то будет выход из цикла
		fmt.Printf("Worker Id %d: %d\n", id, v) //Для наглядности печатаем Worker Id
	}
}

func main(){
	numb := flag.Int("workers", 5, "Number of workers") // go run main.go -workers 10. Defalut:5
	flag.Parse()
	ch := make(chan int)
	var wg sync.WaitGroup
	for i:=0; i<*numb;i++{
		wg.Add(1)
		go Worker(ch, i, &wg)
	}

	// генерация данных
	for id := 0; id < 20; id++ { // ограничиваем количество сообщений для примера
		ch <- id
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) //закрываем канал перед выходом
	wg.Wait() //ждем завершения всех воркеров
} 