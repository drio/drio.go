package main

import (
	"fmt"
)

// isPalindrome checks if a stirng is palindrome
func isPalindrome(s *string) bool {
	l := len(*s)
	for i:=l-1; i>=l/2; i-- {
		if (*s)[(l-1)-i] != (*s)[i] {
			return false
		}
	}
	return true
}

// A palindromic number reads the same both ways. The largest
// palindrome made from the product of two 2-digit numbers is
// 9009 = 91 * 99.
// Find the largest palindrome made from the product of two 3-digit
// numbers.
func main() {
	r := 0
	for i:=1; i<1000; i++ {
		for j:=1; j<1000; j++ {
			p := i*j
			if p > r {
			 	s := fmt.Sprintf("%d", p)
				if isPalindrome(&s) { r = p }
			}
		}
	}
	fmt.Println(r)
}

