package main

import "fmt"

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	m2 := map[string]int{"d": 4, "b": 2, "c": 3, "a": 5}

	if compareMaps(m1, m2) {
		fmt.Println("compare map is valid")
	} else {
		fmt.Println("compare map is not valid")
	}
}

func compareMaps(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		v2, found := m2[k]
		if !found || v1 != v2 {
			return false
		}
	}
	return true
}
