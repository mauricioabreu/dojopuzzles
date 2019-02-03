package main

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
}

// DrawSpiralMatrix draws a multidimensional matrix
// following the example above:
// Given 3 columns and 4 rows
func DrawSpiralMatrix(columns int, rows int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, columns)
	}

	// Populate matrix
	inc := 1
	rowToRight := 0
	columToDown := columns - 1
	rowToLeft := rows - 1
	columnToTop := 0
	goTo := RIGHT
	for {
		if inc >= rows*columns+1 {
			break
		}

		if goTo == RIGHT {
			for c := 0; c < columns; c++ {
				if matrix[rowToRight][c] == 0 {
					matrix[rowToRight][c] = inc
					inc++
				}
			}
			rowToRight++
			goTo = DOWN
		} else if goTo == DOWN {
			for r := 1; r < rows; r++ {
				if matrix[r][columToDown] == 0 {
					matrix[r][columToDown] = inc
					inc++
				}
			}
			columToDown--
			goTo = LEFT
		} else if goTo == LEFT {
			for c := columns - 1; c >= 0; c-- {
				if matrix[rowToLeft][c] == 0 {
					matrix[rowToLeft][c] = inc
					inc++
				}
			}
			rowToLeft--
			goTo = TOP
		} else if goTo == TOP {
			for r := rows - 1; r >= 0; r-- {
				if matrix[r][columnToTop] == 0 {
					matrix[r][columnToTop] = inc
					inc++
				}
			}
			columnToTop++
			goTo = RIGHT
		}
	}

	return matrix
}
