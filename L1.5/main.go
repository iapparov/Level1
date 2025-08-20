package main

import (
	"context"
	"fmt"
	"time"
	"sync"
)

func Worker(wg *sync.WaitGroup, ctx context.Context, r chan int){
	defer wg.Done()
	for {
		select{
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		case val, ok := <-r:
			if ok{
				fmt.Println(val)
			}
		}
	}
}

func main(){

	var N time.Duration
	fmt.Println("Enter time in seconds")
	fmt.Scan(&N)
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel() //обязательно вызываем для освобождения ресурсов
	var wg sync.WaitGroup
	
	wg.Add(1)
	go Worker(&wg, ctx,ch)

	id :=0
	for{
		select{
		case <-ctx.Done():
			fmt.Println("Producer stopped")
			close(ch)
			wg.Wait()
			return
		case ch<-id:
			id++
			time.Sleep(100*time.Millisecond)
		}
	}
} 
/* 
По заданию (подсказке) советуется использовать time.After. Это можно реализовать таким образом, однако здесь приходится создавать еще одну горутину

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	go func(){
		<-time.After(N*time.Second)
		cancel()
	}()

*/