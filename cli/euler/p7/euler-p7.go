package main

import (
  "fmt"
  "github.com/drio/drio.go/math"
)

//    By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can
//    see that the 6th prime is 13.
//
//    What is the 10 001st prime number?
func main() {
  p := 0
  for i := uint64(2); true; i++ {
    if math.IsPrime(i) {
      p += 1
      if p == 10001 {
        fmt.Println(i)
        return
      }
    }
  }
}
