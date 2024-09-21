package main

import "fmt"

func rotateArray(a []int, n int, k int) []int {
	if n <= 0 {
		return nil
	}

	subArr1 := a[:k]
	a = a[k:]
	a = append(a, subArr1...)

	return a
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	n := 5
	k := 2
	subRotate := rotateArray(arr, n, k)
	fmt.Println(subRotate)
}
