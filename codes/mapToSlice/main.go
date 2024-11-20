package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	slice := make([]int, 0, len(m))

	for _, values := range m {
		slice = append(slice, values)
		fmt.Println(slice)
	}
}
