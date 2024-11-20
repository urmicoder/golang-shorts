package main

import "fmt"

func main() {
	m := map[string]int{"a": 2, "b": 1, "c": 4}
	keySlice := make([]string, 0, len(m))

	for key, _ := range m {
		keySlice = append(keySlice, key)
		fmt.Println(keySlice)
	}
}
