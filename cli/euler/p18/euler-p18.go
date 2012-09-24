package main

import (
	"log"
	"fmt"
	"strings"
	"strconv"
	"github.com/drio/drio.go/common/files"
)

// Problem 18
//
//    31 May 2002
//
//    By starting at the top of the triangle below and moving to adjacent
//    numbers on the row below, the maximum total from top to bottom is 23.
//
//    3
//    7 4
//    2 4 6
//    8 5 9 3
//
//    That is, 3 + 7 + 4 + 9 = 23.
//
//    Find the maximum total from top to bottom of the triangle below:
//
//    see file
//
//    NOTE: As there are only 16384 routes, it is possible to solve this
//    problem by trying every route. However, [7]Problem 67, is the same
//    challenge with a triangle containing one-hundred rows; it cannot be
//    solved by brute force, and requires a clever method! ;o)

type triangle struct {
	data [][]int
}

func findMax(t *triangle) int {
	for r:=len(t.data)-1; r>0; r-- {
		for i, v := range t.data[r-1] {
			leftSum  := v + t.data[r][i]
			rightSum := v + t.data[r][i+1]
			if leftSum > rightSum {
				t.data[r-1][i] = leftSum
			} else {
				t.data[r-1][i] = rightSum
			}
		}
	}
	return t.data[0][0]
}

// findMaxRec solves the problem by recursively walking all the
// possible paths starting from the top. cost is O(2^N). It won't scale
// for big triangles
func findMaxRec(row, col, curr int, max *int, t *triangle) {
	//fmt.Println("(", row, ",", col, ") data", t.data[row][col], "curr=", curr, "max=", *max)
	if row == len(t.data) - 1 {
		if curr > *max {
			*max = curr
		}
		return
	}

	couldImprove := ((len(t.data)-row) * 99) + curr > *max
	//couldImprove := true
	if couldImprove {
		pRow, pCol := row, col
		findMaxRec(row+1, col, curr + t.data[row+1][col], max, t)
		row, col = pRow, pCol
		findMaxRec(row+1, col+1, curr + t.data[row+1][col+1], max, t)
	}
}

func load(fn string) *triangle {
	t := new(triangle)
	fd, buff := files.Xopen(fn)
	for line := range files.IterLines(buff) {
	  numbers := []int{}
		for _, n := range strings.Split(line, " ") {
			if i, err := strconv.Atoi(n); err == nil {
				numbers = append(numbers, i)
			} else {
				log.Fatal("Problems parsing line.", err)
			}
		}
		t.data = append(t.data, numbers)
	}
	fd.Close()
	return t
}


func main() {
	//t := load("input.txt")
	//t := load("input-small.txt")
	t := load("triangle.txt")
	//max := 0
	//findMax(0,0,t.data[0][0],&max,t)
	fmt.Println(findMax(t))
}
