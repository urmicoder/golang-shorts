package main

import "fmt"

func div(num int) int {
	if num == 0 {
		fmt.Println("not divisable by zero")
	}
	return 27 / num
}

func rec() {
	r := recover()
	if r != nil {
		fmt.Println("i am recovering from panic")
		fmt.Println("i am fine now")
	}
}

func main() {
	defer rec()
	pa := div(0)
	fmt.Println(pa)
	fmt.Println("main regained")
}
