package main

import (
	"fmt"
	"sync"
)

func producer(m *map[int]string, mu *sync.Mutex, ch chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	vm := *m
	for i := 0; i < 5; i++ {
		mu.Lock()
		vm[i] = fmt.Sprint("$", i)
		mu.Unlock()
		if i == 0 {
			ch <- true
			close(ch)
		}
	}
	// ch <- true
	// close(ch)
	m = &vm
}

func consumer(m *map[int]string, mu *sync.Mutex, ch <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	vm := *m
	<-ch
	for i := 0; i < 5; i++ {
		mu.Lock()
		fmt.Println(vm[i])
		mu.Unlock()
	}
}

func main() {
	m := make(map[int]string)
	mu := sync.Mutex{}
	ch := make(chan bool)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go producer(&m, &mu, ch, &wg)
	go consumer(&m, &mu, ch, &wg)
	wg.Wait()
}
