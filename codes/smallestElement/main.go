package main

import (
	"fmt"
)

func smallestElement(a []int, n int) (min int) {
	if n <= 0 {
		fmt.Println("Array of intger is empty")
		return 0
	}
	min = a[0]

	for i := 0; i < n; i++ { //n
		if a[i] < min { //n
			min = a[i]
		}
	}
	return min

}

// func smallestElement(a []int) int {
// 	if len(a) == 0 {
// 		return 0
// 	}
// 	sort.Ints(a)
// 	return a[0]
// }

func main() {
	arr := []int{9, 5, 8, 1, 3, 6}
	n := 6
	minarr := smallestElement(arr, n)
	//minarr := smallestElement(arr)
	fmt.Println(minarr)
}

//Time Complexity: O(n), where n is the number of elements in the slice.
//Space Complexity: O(1), as the function uses a constant amount of extra space.
