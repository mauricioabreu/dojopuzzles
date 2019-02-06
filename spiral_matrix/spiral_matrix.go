package main

import (
	"fmt"
	"os"
	"strconv"
)

// Direction where to draw
type Direction int

const (
	// RIGHT Draw to right
	RIGHT Direction = iota + 1
	// DOWN Draw to down
	DOWN
	// LEFT Draw to left
	LEFT
	// TOP Draw to top
	TOP
)

func main() {
	columns, _ := strconv.Atoi(os.Args[1])
	rows, _ := strconv.Atoi(os.Args[2])
	spiralMatrix := DrawSpiralMatrix(columns, rows)
	DrawIt(spiralMatrix)
}

// DrawIt outputs the spiral matrix
func DrawIt(spiralMatrix [][]int) {
	for _, r := range spiralMatrix {
		for _, c := range r {
			fmt.Printf("%d ", c)
		}
		fmt.Println()
	}
}

// DrawSpiralMatrix draws a multidimensional matrix
// following the example above:
// Given 3 columns and 4 rows
func DrawSpiralMatrix(columns int, rows int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, columns)
	}

	// Increment every drawn number
	inc := 1

	// These variables are used to know from where to start and where to go
	rowToRight := 0
	columnToRightStart := 0
	columnToRightEnd := columns - 1
	columToDown := columns - 1
	rowToDownStart := 1
	rowToDownEnd := rows - 1
	rowToLeft := rows - 1
	columnToLeftStart := columns - 1
	columnToLeftEnd := 0
	columnToTop := 0
	rowsToTopStart := rows - 2
	rowsToTopEnd := 1

	// Initial direction
	goTo := RIGHT

	// Populate matrix
	for {
		// Stop drawing when we get the last number drawn on the matrix
		if inc >= rows*columns+1 {
			break
		}

		if goTo == RIGHT {
			for c := columnToRightStart; c <= columnToRightEnd; c++ {
				matrix[rowToRight][c] = inc
				inc++
			}
			rowToRight++
			columnToRightStart++
			columnToRightEnd--
			goTo = DOWN
		} else if goTo == DOWN {
			for r := rowToDownStart; r < rowToDownEnd; r++ {
				matrix[r][columToDown] = inc
				inc++
			}
			columToDown--
			rowToDownStart++
			rowToDownEnd--
			goTo = LEFT
		} else if goTo == LEFT {
			for c := columnToLeftStart; c >= columnToLeftEnd; c-- {
				matrix[rowToLeft][c] = inc
				inc++
			}
			rowToLeft--
			columnToLeftStart--
			columnToLeftEnd++
			goTo = TOP
		} else if goTo == TOP {
			for r := rowsToTopStart; r >= rowsToTopEnd; r-- {
				matrix[r][columnToTop] = inc
				inc++
			}
			columnToTop++
			rowsToTopStart--
			rowsToTopEnd++
			goTo = RIGHT
		}
	}

	return matrix
}
