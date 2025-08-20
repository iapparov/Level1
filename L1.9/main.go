package main

import (
	"fmt"
	"sync"
)

func sender(w chan int, array []int, wg *sync.WaitGroup){
	defer wg.Done()
	defer close(w) //Закрываем канал, чтобы не было утечек памяти
	for _, elem := range array{
		w <- elem
	}
}

func reader(r chan int, w chan int, n int, wg *sync.WaitGroup){
	defer wg.Done()
	defer close(r) //Закрываем канал, чтобы не было утечек памяти
	for i := 0; i < n; i++ {
		r <- (<-w) * 2
	}
}

func main(){
	array := []int{1,2,3,4,5,6,7,8} 
	ChIn := make(chan int) //создаем канал для записи
	ChOut := make(chan int) //канал для чтения и обработки результата
	var wg sync.WaitGroup 
	wg.Add(2)
	go sender(ChIn, array, &wg)
	go reader(ChOut, ChIn, len(array), &wg)
	for i := 0; i < len(array); i++ {
		fmt.Println(<-ChOut)
	}
	wg.Wait()
}

/* 

Вообще здесь и без WG все будет работать, т.к. каждый канал будет блокировать свою горутину, и каждая горутина будет анлочить соседнюю. 
Но в этом случае как минимум стоит закрывать каналы через defer, иначе главная горутина может завершиться раньше, чем горутина reader успеет закрыть канал

*/