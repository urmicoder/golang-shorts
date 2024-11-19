package main

import (
	"fmt"
	"sync"
)

func waitcheck(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("i am using waitgroup")
	//defer wg.Done()
	fmt.Println("i am using waitgroup time")
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go waitcheck(&wg)
	wg.Wait()

}
