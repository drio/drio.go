package main

import (
	"fmt"
)

// 2520 is the smallest number that can be divided by each of the
// numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by
// all of the numbers from 1 to 20?
func main() {
	n := 1
	for true {
		for i:=1; i<=20; i++ {
			if n % i != 0 {
				n += 1;
				break
			} else {
				if i == 20 {
					fmt.Println(n)
					return
				}
			}
		}
	}
}
