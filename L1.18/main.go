package main

import (
	"fmt"
	"sync"
    "sync/atomic"
)

type Counter struct{
    num int64
} //По заданимаю требуется реализовать структуру счетчик

func (c *Counter) CounterMutex(mu *sync.Mutex, wg *sync.WaitGroup){
    defer wg.Done()
    for i:=0; i<1000; i++{
        mu.Lock()
        c.num+=1
        mu.Unlock()
    }
} 

func (c *Counter) CounterAtomic(wg *sync.WaitGroup){
    defer wg.Done()
    for i:=0; i<1000; i++{
        atomic.AddInt64(&c.num, 1) //Насколько я понимаю, атомик в этом случае быстрее, так как не требует захвата мьютекса
    }
} 


func main() {
    var numMu, numAtom Counter
    var mu sync.Mutex
    var wg sync.WaitGroup
    for i:=0; i<5; i++{
        wg.Add(2)
        go numMu.CounterMutex(&mu, &wg)
        go numAtom.CounterAtomic(&wg)
    }
    wg.Wait()
    fmt.Println(numMu.num)
    fmt.Println(numAtom.num)
}

