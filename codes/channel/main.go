package main

import (
	"fmt"
	"time"
)

func channelcheck(ch chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go channelcheck(ch)
	for i := range ch {
		<-ch
		fmt.Println(i)
	}
}
