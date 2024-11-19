package main

import "fmt"

func sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

func median(a []int, n int) int {
	arr := sort(a)
	if n <= 0 {
		return 0
	}
	mid := 0
	average := 0
	mid1 := arr[n/2-1]
	mid2 := arr[n/2]
	average = (mid1 + mid2) / 2
	if n%2 == 0 {
		mid = average
	} else {
		mid = arr[n/2]
	}

	return mid
}

func main() {
	arr := []int{2, 5, 1}
	n := len(arr)
	midArray := median(arr, n)
	fmt.Println(midArray)
}
