package main

import "fmt"

func secondMaxMinElement(a []int, n int) (secondMax int, secondMin int) {
	if n <= 0 {
		return 0, 0
	}

	min := a[0]
	max := a[0]
	secondMin = a[0]
	secondMax = a[0]

	for i := 0; i < n; i++ {
		if a[i] < min {
			secondMin = min
			min = a[i]
		} else {
			if a[i] < secondMin && a[i] != min {
				secondMin = a[i]
			}
		}
	}

	for i := 0; i < n; i++ {
		if a[i] > max {
			secondMax = max
			max = a[i]
		} else {
			if a[i] > secondMax && a[i] != max {
				secondMax = a[i]
			}
		}
	}

	return secondMax, secondMin
}

func main() {
	arr := []int{2, 5, 8, 1, 3, 6, 7, 0}
	n := 8
	elementMax, elementMin := secondMaxMinElement(arr, n)
	fmt.Println(elementMax, elementMin)
}
