package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "d", "b", "a"}
	map1 := make(map[string]int)

	for _, k := range slice {
		map1[k]++
	}
	fmt.Println(map1)
}
