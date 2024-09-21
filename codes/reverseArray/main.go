package main

import "fmt"

// func reverseArray(a []int, n int) (result []int) {
// 	if n <= 0 {
// 		return nil
// 	}
// 	for i := 0; i < n/2; i++ {
// 		temp := a[i]
// 		a[i] = a[n-i-1]
// 		a[n-i-1] = temp
// 	}
// 	return a
// }

func reverseArray(a []int, n int) (result []int) {
	if n <= 0 {
		return nil
	}
	l := 0
	r := n - 1
	for l < r {
		temp := a[l]
		a[l] = a[r]
		a[r] = temp
		l++
		r--
	}
	return a
}

func main() {
	arr := []int{5, 4, 3, 2, 1, 6}
	n := 6
	reverse := reverseArray(arr, n)
	fmt.Println(reverse)
}
