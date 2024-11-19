package main

import "fmt"

// closer function is a special type of anonomus function that can access variabeles which is declerd ouside of the function.
// closer func allow us to difined fuctions to variables, it use function as a argument of a func or as a return value of function
// here is a example
func main() {
	v := func(f func(i int) int) func() int {
		i := f(3)
		return func() int {
			//i++
			fmt.Println("here is a closer function")
			return i
		}
	}
	f := v(func(i int) int {
		return i * 2
	})
	retur := f()
	fmt.Println(retur)
}
