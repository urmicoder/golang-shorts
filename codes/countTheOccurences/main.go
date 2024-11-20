package main

import "fmt"

func main() {
	m := make(map[string]int)
	arrayOfString := []string{"a", "b", "a", "c", "b", "c", "a", "c", "c", "a", "b"}

	for _, value := range arrayOfString {
		m[value]++
		//fmt.Println(m)
	}
	fmt.Println(m)
}
