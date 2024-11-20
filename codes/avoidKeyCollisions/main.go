package main

import "fmt"

func main() {
	m1 := map[string]int{"a": 1, "b": 3, "c": 4, "d": 7}
	m2 := map[string]int{"x": 3, "y": 6, "z": 8, "w": 5} //need to use uniqes key name or using prefix
	combineMap := make(map[string]int)

	for k, v := range m1 {
		combineMap[k] = v
	}

	for k, v := range m2 {
		combineMap[k] = v
	}
	fmt.Println(combineMap)
}
