package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"a": 8, "d": 5, "r": 1, "p": 9, "c": 3}
	arrayString := make([]int, 0, len(m))

	for _, value := range m {
		arrayString = append(arrayString, value)
		fmt.Println(arrayString)
	}
	sort.Ints(arrayString)
	fmt.Println(arrayString)

	for _, v := range arrayString { //jab index ke jagha undercor nhi tha correct resp. nhi mil raha tha
		for k, val := range m {
			if v == val {
				fmt.Println(k, v)
				break
			}
		}
	}
}
