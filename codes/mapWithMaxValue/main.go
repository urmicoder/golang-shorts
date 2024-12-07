package main

import (
	"fmt"
	"math"
)

// finding the max value in map
func main() {
	m := map[string]int{"a": 1, "b": 8, "c": 3, "d": 5}

	maxValue := math.MinInt
	var maxKey string

	for k, v := range m {
		if v > maxValue {
			maxValue = v
			maxKey = k
			//fmt.Println(maxKey, maxValue)
		}
	}
	//fmt.Println(maxValue)
	fmt.Println(maxKey, maxValue)
}
