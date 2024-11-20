package main

import "fmt"

func main() {
	type Person struct {
		FirstName string
		LastName  string
	}

	mp := make(map[Person]int)
	//p := Person{FirstName: "Urmi", LastName: "Kewat"}
	p := Person{"Urmi", "Kewat"}
	mp[p] = 27
	fmt.Println(mp)
}
