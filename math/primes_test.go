package math

import "testing"

import (
)

func check(truth, result []uint64, t *testing.T) {
	for i:=0; i<len(truth); i++ {
		if result[i] != truth[i] {
			t.Errorf("PrimeSieve(30) failed")
		}
	}
}

func TestPrimes(t *testing.T) {
	truth := []uint64{}
	check(truth, PrimeSieve(0), t)
  truth = []uint64{2, 3, 5, 7}
	check(truth, PrimeSieve(10), t)
  truth = []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	check(truth, PrimeSieve(30), t)
}

func TestPrimeFactorsOf(t *testing.T) {
	truth := []uint64{5, 5}
	check(truth, PrimeFactorsOf(25), t)
}
