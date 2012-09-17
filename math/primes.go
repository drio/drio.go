package math

import "math"

//import "fmt"

// PrimeSieve finds all prime numbers up to n.
// Uses the Sieve of Eratosthenes algorithm without optimizations
// TODO: Add optimizations
func PrimeSieve(n uint64) []uint64 {
  isPrime := make([]bool, n)
  var i uint64
  for i = 2; i < n; i++ {
    isPrime[i] = true
  }

  limit := uint64(math.Sqrt(float64(n))) + 1
  for i = 2; i < limit; i++ {
    if isPrime[i] {
      for j := i + 1; j < n; j++ {
        if j%i == 0 {
          isPrime[j] = false
        }
      }
    }
  }

  primes := []uint64{}
  for i = 1; i < n; i++ {
    if isPrime[i] {
      primes = append(primes, i)
    }
  }

  return primes
}

// PrimeFactorsOf returns a slice with the factors of the number n
func PrimeFactorsOf(n uint64) []uint64 {
  if n == 1 {
    return []uint64{1}
  }
  primes := PrimeSieve(uint64(math.Sqrt(float64(n)) + 1))
  primeFactors := []uint64{}

  for _, p := range primes {
    if p*p > n {
      break
    }
    for n%p == 0 {
      primeFactors = append(primeFactors, p)
      n = n / p
    }
  }
  if n > 1 {
    primeFactors = append(primeFactors, n)
  }

  return primeFactors
}

// IsPrime tells you if a number is prime or not.
func IsPrime(n uint64) bool {
  pf := PrimeFactorsOf(n)
  if len(pf) == 1 && pf[0] == n {
    return true
  }
  return false
}
