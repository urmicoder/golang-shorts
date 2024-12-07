package main

import (
	"fmt"
	"net"
)

func isValid(ip string) bool {
	return net.ParseIP(ip) != nil
}

func main() {
	ip := "255.56.108.0"
	//ip := "hello"

	if isValid(ip) {
		fmt.Println(ip, "is a valid IP Address")
	} else {
		fmt.Println(ip, "is not a valid IP Address")
	}
}
