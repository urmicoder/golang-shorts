package main

import (
	"fmt"
	"math"
)

func main() {
	m := map[string]int{"a": 1, "b": 8, "c": 3, "d": 5}

	minValue := math.MaxInt
	var minKey string
	for key, value := range m {
		if value < minValue {
			minValue = value
			minKey = key
		}
	}
	fmt.Println(minKey, minValue)
}
