package main

import (
	"fmt"
)

func averageArray(a []int, n int) int {
	if n <= 0 {
		return 0
	}
	sum := 0
	average := 0

	for _, value := range a {
		sum += value
		average = sum / n
	}
	return average
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := 5
	aveElement := averageArray(arr, n)
	fmt.Println(aveElement)
}
