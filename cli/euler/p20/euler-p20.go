package main

import (
  "fmt"
  "math/big"
  "strconv"
)

// Problem 20
//
//    21 June 2002
//
//    Find the sum of the digits in the number 100!
func main() {
  fact := big.NewInt(1)
  for i := 2; i <= 100; i++ {
    bi := big.NewInt(int64(i))
    fact.Mul(fact, bi)
  }

  bs := fact.String()
  sum := 0
  for i := 0; i < len(bs); i++ {
    v, _ := strconv.Atoi(bs[i : i+1])
    sum += v
  }
  fmt.Println(bs, sum)
}
