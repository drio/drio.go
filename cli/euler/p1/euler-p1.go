package main

import (
  "fmt"
)

// Add all the natural numbers below one thousand that are multiples of 3 or 5.
func main() {
  sum := 0
  for i := 1; i < 1000; i++ {
    if i%3 == 0 || i%5 == 0 {
      sum += i
    }
  }
  fmt.Println(sum)
}
