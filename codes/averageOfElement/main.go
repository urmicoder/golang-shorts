package main

import (
	"fmt"
	"sync"
)

// func averageArray(a []int, n int) int {
// 	if n <= 0 {
// 		return 0
// 	}
// 	sum := 0
// 	average := 0

// 	for _, value := range a {
// 		sum += value
// 		average = sum / n
// 	}
// 	return average
// }

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	n := 5
// 	aveElement := averageArray(arr, n)
// 	fmt.Println(aveElement)
// }

func calculateSum(a []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	sum := 0
	for _, value := range a {
		sum += value
	}

	ch <- sum // Send the sum through the channel
}

func averageArray(a []int) int {
	// Use a WaitGroup to wait for the goroutine to finish
	var wg sync.WaitGroup

	// Create a channel to communicate the sum
	ch := make(chan int)

	// Start the goroutine to calculate the sum
	wg.Add(1)
	go calculateSum(a, ch, &wg)

	// Wait for the goroutine to finish
	wg.Wait()

	// Receive the sum from the channel
	sum := <-ch

	// Calculate the average
	average := sum / len(a)
	return average
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	aveElement := averageArray(arr)
	fmt.Println(aveElement) // Output: 3
}
