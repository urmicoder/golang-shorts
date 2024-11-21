package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"a": 1, "r": 3, "k": 8, "c": 5, "b": 2, "h": 7}
	arrayString := make([]string, 0, len(m))

	for k, _ := range m {
		arrayString = append(arrayString, k)
		fmt.Println(arrayString)
	}
	sort.Strings(arrayString)

	for _, k := range arrayString {
		fmt.Println(k, m[k])
	}
}
