package main

import "fmt"
import "github.com/drio/drio.go/math"

// Problem 10
//
//    08 February 2002
//
//    The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
//    Find the sum of all the primes below two million.
func main() {
  sum := uint64(0)
  for _, e := range math.PrimeSieve(2000000) {
    sum += e
  }
  fmt.Println(sum)
}
