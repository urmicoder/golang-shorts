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

func rearrangeArray(a []int, n int) []int {
	if n <= 0 {
		return nil
	}
	mid := (n / 2)
	value := (n + mid) / 2
	for i := mid; i < value; i++ {
		t := a[i]
		j := n - 1 - (i - mid)
		a[i] = a[j]
		a[j] = t
	}
	return a
}

func main() {
	arr := []int{4, 2, 8, 6, 15, 5, 9, 20}
	n := 7
	sortArr := sort(arr)
	fmt.Println(sortArr)
	rearr := rearrangeArray(sortArr, n)
	fmt.Println(rearr)
}
