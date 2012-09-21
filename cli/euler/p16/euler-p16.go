package main

import (
  "fmt"
  "math/big"
  "strconv"
)

// Problem 16
//
//    03 May 2002
//
//    2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//
//    What is the sum of the digits of the number 2^1000?
func main() {
  b := big.NewInt(1)
  two := big.NewInt(2)
  for i := 0; i < 1000; i++ {
    b.Mul(b, two)
  }
  bs := b.String()
  sum := 0
  for i := 0; i < len(bs); i++ {
    v, _ := strconv.Atoi(bs[i : i+1])
    sum += v
  }
  fmt.Println(sum)
}
