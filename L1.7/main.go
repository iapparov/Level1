package main

import (
	"sync"
)

func Worker1(wg *sync.WaitGroup, mu *sync.Mutex, storageStd map[int]string, idx int){
	defer wg.Done()
	for i:= idx; i<idx+10; i++{
		mu.Lock()
		storageStd[i]="Test"
		mu.Unlock()
	}
} 

func Worker2(wg *sync.WaitGroup, storageStd *sync.Map, idx int){
	defer wg.Done()
	for i:= idx; i<idx+10; i++{
		storageStd.Store(i, "Test")
	}
} 


func main(){

	var wg sync.WaitGroup
	var mu sync.Mutex // Используем мьютекс, для избежания гонок с обычной map. 
	var storageSync sync.Map //Это мапа подходит для конкуретной работы из коробки, не нужно использовать мьютексы. 
	storageStd := make(map[int]string, 50) //Выделим память сразу под 50 элементов

	for i := range 5{
		wg.Add(2)
		go Worker1(&wg, &mu, storageStd, i)
		go Worker2(&wg, &storageSync, i)
	}

	wg.Wait()
}

/* 
Вместо sync.Mutex можно было бы использовать sync.RWMutex, но насколько я знаю, он влияет на скорость, когда много чтения -> можно юзать mu.RLock/mu.RUnlock
*/