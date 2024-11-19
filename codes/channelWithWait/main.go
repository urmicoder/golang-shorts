package main

import (
	"fmt"
	"sync"
	"time"
)

func channelcheck(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second) //yiske bina reponse bina time liye mil jata hai
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go channelcheck(ch, &wg)
	for i := range ch {
		<-ch
		fmt.Println(i)
	}
	wg.Wait()
}
