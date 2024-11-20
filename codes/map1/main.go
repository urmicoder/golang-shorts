package main

import "fmt"

func main() {
	mp := make(map[string]int)
	mp["urmi"] = 27
	mp["rupam"] = 29
	mp["pam"] = 28
	fmt.Println(mp)

	value := mp["rupam"]
	fmt.Println(value)
	value, exit := mp["chi"]
	if exit {
		fmt.Println("urmi value exit:", value)
	} else {
		fmt.Println("urmi value not found")
	}
	delete(mp, "pam")
	fmt.Println(mp)
}
