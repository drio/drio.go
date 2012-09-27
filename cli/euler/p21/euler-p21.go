package main

import (
  "fmt"
)

// Problem 21
//
//    05 July 2002
//
//    Let d(n) be defined as the sum of proper divisors of n (numbers less
//    than n which divide evenly into n).
//    If d(a) = b and d(b) = a, where a != b, then a and b are an amicable
//    pair and each of a and b are called amicable numbers.
//
//    For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22,
//    44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are
//    1, 2, 4, 71 and 142; so d(284) = 220.
//
//    Evaluate the sum of all the amicable numbers under 10000.
const upto = 10000

func findSum() map[int]int {
  md := make(map[int]int) // map of d() computation results
  for i := upto; i > 0; i-- {
    for j := 1; j < i; j++ {
      if i%j == 0 {
        md[i] += j
      }
    }
  }
  return md
}

func main() {
  mSum := make(map[int]int)
  mDivs := findSum()
  for num, sum := range mDivs {
    if num != sum && num == mDivs[mDivs[num]] {
      mSum[num]++
      mSum[sum]++
    }
  }
  sumAmicable := 0
  for k, _ := range mSum {
    fmt.Println(k)
    sumAmicable += k
  }
  fmt.Println("> ", sumAmicable)
}
