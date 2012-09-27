package math

import (
  "math/big"
  "testing"
)

func TestSum(t *testing.T) {
  if 6 != SumDigits(big.NewInt(123)) {
    t.Errorf("Error adding digits of number 123")
  }
  if 0 != SumDigits(big.NewInt(0)) {
    t.Errorf("Error adding digits of number 0")
  }
  if 15 != SumDigits(big.NewInt(735)) {
    t.Errorf("Error adding digits of number 735")
  }
}
