package main

import "fmt"

func main() {
	nastedMap := make(map[string]map[string]int)
	nastedMap["urmi"] = make(map[string]int)
	// nastedMap["urmi"]["x"] = 2
	// nastedMap["rupam"]["y"] = 3 aise chije nhi hoti mujhe yisme ek aur map banana hoga tab ye cooect hoga
	// nastedMap["pam"]["y"] = 7
	nastedMap["urmi"]["x"] = 1
	nastedMap["urmi"]["y"] = 7
	nastedMap["urmi"]["z"] = 8
	nastedMap["rupam"] = make(map[string]int)
	nastedMap["rupam"]["a"] = 6
	nastedMap["rupam"]["b"] = 4
	nastedMap["rupam"]["c"] = 5
	fmt.Println(nastedMap)
}
