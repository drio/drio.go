package math

import (
  "math/big"
  "strconv"
)

// SumDigits returns the sum of the digits in bNumber
func SumDigits(bNumber *big.Int) int {
  bs := bNumber.String()
  sum := 0
  for i := 0; i < len(bs); i++ {
    v, _ := strconv.Atoi(bs[i : i+1])
    sum += v
  }
  return sum
}
