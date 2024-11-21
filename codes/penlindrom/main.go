package main

import "fmt"

func isPalindrome(n int) bool {
	if n <= 0 {
		return false
	}
	orginal := n
	reverse := 0
	for n > 0 {
		digit := n % 10
		reverse = reverse*10 + digit
		n = n / 10
	}
	return orginal == reverse
}

func main() {
	//number := []int{121, 1441, 131, 1555, 1671}
	numbers := []int{121, 123, 12321, 456, 909}

	for _, num := range numbers {
		// if isPalindrome(num) {
		// 	fmt.Printf("%d is a palindrome\n", num)
		// } else {
		// 	fmt.Printf("%d is not a palindrome\n", num)
		// }
		fmt.Println(isPalindrome(num))
	}
}
