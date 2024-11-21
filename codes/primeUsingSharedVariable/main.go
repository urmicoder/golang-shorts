package main

import (
	"fmt"
	"sync"
	"time"
)

func isPrime(n int) bool {
	if n <= 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeHelper(a []int, f func(n int) bool, m *map[int]bool, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	vm := *m

	for i := range a {
		mu.Lock()
		vm[i] = f(a[i])
		mu.Unlock()
	}
	time.Sleep(time.Second)
	m = &vm
}

func main() {
	m := make(map[int]bool)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	arr := arrGen()
	nogou := 6
	parts := (len(arr) + nogou - 1) / nogou

	for i := 0; i < nogou; i++ {
		wg.Add(1)
		s := parts * i
		e := s + parts
		if e > len(arr) {
			e = len(arr)
		}
		go primeHelper(arr[s:e], isPrime, &m, &mu, &wg)
	}
	wg.Wait()
	fmt.Println(m)
}

func arrGen() []int {
	var arr []int
	for i := 0; i < 100; i++ {
		arr = append(arr, i)
	}
	return arr
}
