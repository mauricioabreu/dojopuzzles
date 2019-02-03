package main

import (
	"reflect"
	"testing"
)

// 1 2
// 4 3
func TestSpiralMatrix2x2(t *testing.T) {
	columns := 2
	rows := 2
	spiralMatrix := DrawSpiralMatrix(columns, rows)
	if !reflect.DeepEqual(spiralMatrix, [][]int{{1, 2}, {4, 3}}) {
		t.Error("Wrong spiral format")
	}
}

// 1 2
// 6 3
// 5 4
func TestSpiralMatrix2x3(t *testing.T) {
	columns := 2
	rows := 3
	spiralMatrix := DrawSpiralMatrix(columns, rows)
	if !reflect.DeepEqual(spiralMatrix, [][]int{{1, 2}, {6, 3}, {5, 4}}) {
		t.Error("Wrong spiral format")
	}
}

//  1  2  3  4  5
// 18 19 20 21  6
// 17 28 29 22  7
// 16 27 30 23  8
// 15 26 25 24  9
// 14 13 12 11 10
func TestSpiralMatrix5x6(t *testing.T) {
	columns := 5
	rows := 6
	spiralMatrix := DrawSpiralMatrix(columns, rows)
	if !reflect.DeepEqual(spiralMatrix, [][]int{{1, 2, 3, 4, 5}, {18, 19, 20, 21, 6}, {17, 28, 29, 22, 7}, {16, 27, 30, 23, 8}, {15, 26, 25, 24, 9}, {14, 13, 12, 11, 10}}) {
		t.Error("Wrong spiral format")
	}
}
