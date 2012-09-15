package main

import (
	"fmt"
)

// fibIter is an iterator that yields the terms in the
// fibonacci sequence. It stops when reaches fib terms
// bigger than max
func fibIter(max int, ch chan int) {
	prev := make([]int, 2)
	for i:=0; prev[0]+prev[1]<=max ; i++ {
		if i < 2 {
			ch <- i
		} else if i == 2 {
			ch <- 1
			prev[0] = 1; prev[1] = 1
		} else {
			ch <- prev[0] + prev[1]
			prev[0], prev[1] = prev[1], prev[0] + prev[1]
		}
	}
	close(ch)
}

// By considering the terms in the Fibonacci sequence whose
// values do not exceed four million, find the sum of the even-valued
// terms.
func main() {
	ch := make(chan int, 10000)
	go fibIter(4000000, ch)
	sum := 0
	for t:= range ch {
		if t % 2 == 0 {
			sum += t
		}
	}
	fmt.Println(sum)
}

