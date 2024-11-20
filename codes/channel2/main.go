package main

import (
	"fmt"
	"sync"
	"time"
)

func channelW(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	close(ch)
}

func channelR(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
		//<-ch
	}

}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go channelW(ch, &wg)
	go channelR(ch, &wg)
	wg.Wait()
}
