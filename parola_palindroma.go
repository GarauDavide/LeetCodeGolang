package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1188))
	fmt.Println(isPalindrome(1234321))
	fmt.Println(isPalindrome(123321))
	fmt.Println(isPalindrome(-222))
	fmt.Println(isPalindrome(2442))
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	strLen := len(str)
	forCondition := (strLen / 2) + 1

	for i := 0; i < forCondition; i++ {

		if str[i] != str[strLen-(i+1)] {
			return false
		}
	}

	return true
}
