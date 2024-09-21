package main

import "fmt"

func swap(a, b int) (int, int) {
	//a = a + b
	//b = a - b
	//a = a - b
	return b, a
}

func main() {
	var a, b int = 8, 9
	snum, snum2 := swap(a, b)
	fmt.Println(snum, snum2)
}
