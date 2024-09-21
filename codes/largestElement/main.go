package main

import "fmt"

func largestElement(a []int, n int) (max int) {
	if n <= 0 {
		fmt.Println("Array is empty")
		return 0
	}
	max = a[0]
	for i := 0; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func main() {
	arr := []int{2, 5, 8, 3, 1, 0}
	n := 5
	maxArr := largestElement(arr, n)
	fmt.Println(maxArr)
}
