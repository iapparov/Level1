package main

import (
	"fmt"
	"time"
)


func Sleep1(i time.Duration){ //Просто слипаем обычным циклом
	t := time.Now()
	for {
		if time.Since(t) >=i{ 
			break
		}
	}
}

func Sleep2(i time.Duration){
	<-time.After(i) //блокируем чтение из канала пока туда не прилетит значение через time.After
}

func Sleep3(i time.Duration){
	ch := make(chan struct{})
	go func(){
		time.AfterFunc(i, func() {
			ch <- struct{}{}
		}) //юзаем колбэк через i времени
	}()
	<- ch //блокируем горутину чтением из канала
}

func main() {
	fmt.Println("Start: ", time.Now())
	Sleep1(1*time.Second)
	Sleep2(1*time.Second)
	Sleep3(1*time.Second)
	fmt.Println("End: ", time.Now())
}