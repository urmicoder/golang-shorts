package main

import (
	"fmt"
)

func countFrequency(a []int, n int) map[int]int {
	if n <= 0 {
		return nil
	}
	element := make(map[int]int)

	for _, value := range a {
		element[value]++
	}
	return element
}

func main() {
	arr := []int{10, 5, 10, 15, 10, 5}
	n := len(arr)
	count := countFrequency(arr, n)
	for key, value := range count {
		fmt.Println(key, value)
	}
	//fmt.Println(count)
}
