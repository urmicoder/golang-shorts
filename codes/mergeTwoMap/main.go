package main

import "fmt"

// merging two map me simplaly dono map se key and value nikal kar usko new map me key ke crossponding values dal dene hai.
// To merge two maps simply, we need to extract the keys and values from both maps and insert them into a new map, placing the values corresponding to each key.
func main() {
	m1 := map[string]int{"a": 1, "b": 8, "c": 3, "d": 5}
	m2 := map[string]int{"e": 1, "f": 8, "c": 3}

	merge := make(map[string]int)
	for k, v1 := range m1 {
		merge[k] = v1
	}
	for k, v2 := range m2 {
		merge[k] = v2
	}
	fmt.Println("created new map after merging:", merge)

}
