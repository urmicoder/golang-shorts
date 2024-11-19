package main

import (
	"fmt"
	"reflect"
)

func main() {
	inter("hello")

}

func inter(i interface{}) {
	// switch i.(type) {
	// case int:
	// 	fmt.Println("integer")
	// case string:
	// 	fmt.Println("Sring hai")
	// default:
	// 	fmt.Println("any type")
	// }
	res := reflect.TypeOf(i)
	fmt.Println(res)
}
