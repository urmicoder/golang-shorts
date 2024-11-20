package main

import (
	"fmt"
	"sync"
	"time"
)

func odder(n int, ch1 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	if n%2 != 0 {
		time.Sleep(time.Second)
		fmt.Println("number is odd")
		ch1 <- n
	}
}

func evener(n int, ch2 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	if n%2 == 0 {
		fmt.Println("number is even")
		ch2 <- n
	}
}

func main() {
	intArray := []int{2, 26, 68, 43, 98, 1, 6, 5, 11}
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := sync.WaitGroup{}
	//wg.Add(2)
	for _, value := range intArray {
		wg.Add(2)
		go odder(value, ch1, &wg)
		go evener(value, ch2, &wg)
	}

	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()

	go func() {
		for odd := range ch1 {
			fmt.Println(odd)
		}
	}()

	go func() {
		for even := range ch2 {
			fmt.Println(even)
		}
	}()
	wg.Wait()
}
