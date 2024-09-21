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
	sum := 0
	average := 0
	for _, value := range arr {
		sum += value
		average = sum / 2
		if n%2 == 0 {
			mid = average
		} else {
			mid = n / 2
		}
	}

	return mid
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := 5
	midArray := median(arr, n)
	fmt.Println(midArray)
}
