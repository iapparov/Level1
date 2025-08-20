package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Worker1(flag *bool, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		fmt.Println("Worker 1 is up")
		time.Sleep(1*time.Second)
		if *flag {
			fmt.Println("Stopped Worker with flag")
			return
		}
	}
} //Завершение по флагу

func Worker2(ctx context.Context, wg *sync.WaitGroup){
	defer wg.Done()
	for{
		select{
			case <-ctx.Done():
				fmt.Println("Stopped Worker with context")
				return
			default:
				time.Sleep(1*time.Second)
				fmt.Println("Worker 2 is up")
		}
	}
}//Завершение по контексту

func Worker3(done chan any, wg *sync.WaitGroup){
	defer wg.Done()
	for{
		select{
		case <-done:
			fmt.Println("Stopped Worker with channel")
			return
		default:
			time.Sleep(1*time.Second)
			fmt.Println("Worker 3 is up")
		}
	}
}//Завершение по каналу

func Worker4(wg *sync.WaitGroup){
	defer func() {
		fmt.Println("Stopped worker with runtime.Goexit()")
		time.Sleep(10 * time.Millisecond) // небольшой буфер для вывода Stopped worker with runtime.Goexit(). Если defer fmt.Println("Stopped worker with runtime.Goexit()") defer wg.Done() то вывод будет через раз
		wg.Done()
	}()
	fmt.Println("Worker 4 is up")
	time.Sleep(1*time.Second)
	runtime.Goexit()
}//Завершение по Goexit


func main(){

	var wg sync.WaitGroup
	wg.Add(4)

	flag := true
	go Worker1(&flag, &wg)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go Worker2(ctx, &wg)

	done := make(chan any)
	go Worker3(done, &wg)
	time.Sleep(1*time.Second)
	close(done)

	go Worker4(&wg)

	wg.Wait()
}

/*
Хардкорное решение, но тоже по сути останавливает горутину

func Worker5(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Worker 5 is up")
	time.Sleep(1*time.Second)
	panic("Stopped Worker with panic")
}

*/