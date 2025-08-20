package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"sync"
)

func Worker(wg *sync.WaitGroup, ctx context.Context, r chan int){ //Принимаем wg по указателю, чтобы не копировать экземпляры
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

	numb := flag.Int("workers", 1, "Number of workers")
	flag.Parse()
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background()) //Создаем контектст (интерфейс для общения между горутинами) с отменой, в качестве аргумента передаем parent (корневой) context
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT)
	var wg sync.WaitGroup
	
	go func(){
		<-sigch //По суми блокируем горутину до момент записи в канал SIGINT
		fmt.Println("Received interrupt signal, shutting down...")
		cancel()
	}()
	
	for i:=0; i<*numb;i++{
		wg.Add(1)
		go Worker(&wg, ctx,ch) 
	}
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
В целом можно сделать обработку SIGINT через канал, а не через контекст. Но через контекст идиоматичнее 
Просто создаем канал done := make(chan any) (any = interface{})
ждем <-sigch и после этого закрываем канал close(done)

в селектах вместо <-ctx.Done() делаем <-done (так как закрытие канала снимает блокировку на чтение)
*/