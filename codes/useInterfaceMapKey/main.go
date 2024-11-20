package main

import "fmt"

func main() {
	type key interface{}
	m := make(map[key]int)
	m["urmi"] = 3
	m["rupam"] = 4
	m[1] = 5
	m[true] = 7
	fmt.Println(m)
}
