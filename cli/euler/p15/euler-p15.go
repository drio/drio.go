package main

import "fmt"

// Problem 15
//
//    19 April 2002
//
//    Starting in the top left corner of a 2 x 2 grid, there are 6 routes
//    (without backtracking) to the bottom right corner.
//
//    How many routes are there through a 20x20 grid?
type vertex struct {
	row, col int
}

func (v *vertex) Forward(gridSize int) bool {
	if v.col == gridSize {
		return false
	}
	v.col += 1
	return true
}

func (v *vertex) Down(gridSize int) bool {
	if v.row == gridSize {
		return false
	}
	v.row += 1
	return true
}

func (v *vertex) Up() bool {
	if v.row == 0 {
		return false
	}
	v.row -= 1
	return true
}

func (v *vertex) Back() bool {
	if v.row == 0 {
		return false
	}
	v.col -= 1
	return true
}

func (v *vertex) isSolution(gridSize int) bool {
	return gridSize == v.row && gridSize == v.col
}

func findNumRoutes(gridSize int, cv *vertex, m map[vertex]uint64) uint64{
	if cv.isSolution(gridSize) {
		return 1
	}

	var sDown, sForward uint64
	if cv.Down(gridSize) {
		if v, present := m[*cv]; present {
			sDown = v
		} else {
			sDown  = findNumRoutes(gridSize, cv, m)
			m[*cv] = sDown
		}
		cv.Up()
	}
	if cv.Forward(gridSize) {
		if v, present := m[*cv]; present {
			sForward = v
		} else {
			sForward  = findNumRoutes(gridSize, cv, m)
			m[*cv] = sForward
		}
		cv.Back()
	}
	return sDown + sForward
}

func main() {
	m := make(map[vertex]uint64)
	fmt.Println(findNumRoutes(20, &vertex{0,0}, m))
}
