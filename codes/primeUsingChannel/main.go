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

func primeHelper(a []int, f func(n int) bool, ch chan map[int]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	m := make(map[int]bool)
	//ch <- m
	for i := range a {
		m[i] = f(i)
	}
	time.Sleep(time.Second)
	ch <- m
}

func main() {
	arr := arrGen()
	nogou := 6
	parts := (len(arr) + nogou - 1) / nogou
	ch := make(chan map[int]bool, nogou)
	m := make(map[int]bool)
	wg := sync.WaitGroup{}
	for i := 0; i < nogou; i++ {
		wg.Add(1)
		s := parts * i
		e := s + parts
		if e > len(arr) {
			e = len(arr)
		}
		go primeHelper(arr[s:e], isPrime, ch, &wg)
	}
	wg.Wait()
	close(ch)
	for i := range ch {
		for k, b := range i {
			m[k] = b
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, ":", m[i])
	}
}

func arrGen() []int {
	var arr []int

	for i := 0; i < 100; i++ {
		arr = append(arr, i)
	}
	return arr
}
