package main

import (
	"fmt"
  "github.com/drio/drio.go/math"
)


// What is the largest prime factor of the number 600851475143
func main() {
	max := uint64(0)
	for _, e := range math.PrimeFactorsOf(600851475143) {
		if e > max { max = e }
	}
	fmt.Println(max)
}

