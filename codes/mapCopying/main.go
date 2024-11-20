package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["ur"] = 2
	m1["urmi"] = 3
	fmt.Println("data:", m1)

	m2 := make(map[string]int)
	fmt.Println("without data:", m2)

	for key, value := range m1 {
		m2[key] = value
	}
	fmt.Println("withdata:", m2)
}
