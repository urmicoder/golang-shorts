package main

import "fmt"

func sumElement(a []int, n int) int {
	if n <= 0 {
		return 0
	}
	sum := 0
	for _, value := range a {
		sum += value
	}
	return sum
}
func main() {
	arr := []int{1, 2, 1, 1, 5, 1}
	n := 5
	sumArray := sumElement(arr, n)
	fmt.Println(sumArray)
}
