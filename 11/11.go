package main

import (
	"fmt"
	"helper"
	"strings"
)

// -----------------------------------------------------------------

func getKdx(idx int, jdx int, cols int) int {
	return idx*cols + jdx
}

func dump(rows int, cols int, matrix *string) {
	fmt.Println(strings.Repeat("-", cols))
	for idx := 0; idx < rows; idx++ {
		for jdx := 0; jdx < cols; jdx++ {
			kdx := getKdx(idx, jdx, cols)
			fmt.Printf("%c ", (*matrix)[kdx])
		}
		fmt.Printf("\n")
	}
	fmt.Println(strings.Repeat("-", cols))
}

func occupy(x int, y int, rows int, cols int, matrix *string) int {
	seats := 0

	for idx := x - 1; idx <= x+1; idx++ {
		if idx < 0 || idx >= rows {
			continue
		}
		for jdx := y - 1; jdx <= y+1; jdx++ {
			if jdx < 0 || jdx >= cols {
				continue
			}
			if x == idx && y == jdx {
				continue
			}
			kdx := getKdx(idx, jdx, cols)
			if (*matrix)[kdx] == '#' {
				seats++
			}
		}
	}
	//fmt.Println("busy seats around ", x, " - ", y, " -> ", seats)
	return seats
}

func parseOne(rows int, cols int, matrix string) {

	oldmatrix := ""
	for matrix != oldmatrix {
		oldmatrix = matrix
		matrix = ""
		for idx := 0; idx < rows; idx++ {
			for jdx := 0; jdx < cols; jdx++ {
				kdx := getKdx(idx, jdx, cols)
				seat := oldmatrix[kdx]
				busy := occupy(idx, jdx, rows, cols, &oldmatrix)
				if seat != '.' && busy == 0 {
					matrix += "#"
				} else if seat != '.' && busy >= 4 {
					matrix += "L"
				} else {
					matrix += string(seat)
				}
			}
		}
		// dump(rows, cols, &matrix)
	}

	dump(rows, cols, &matrix)
	fmt.Println("parseOne - occupied seats ", strings.Count(matrix, "#"))

}

// -----------------------------------------------------------------

// https://www.geeksforgeeks.org/program-to-print-the-diagonals-of-a-matrix/
// https://www.geeksforgeeks.org/n-queen-problem-backtracking-3/?ref=lbp
// https://www.geeksforgeeks.org/print-matrix-diagonal-pattern/?ref=lbp

func drawSlopesX(x int, y int, rows int, cols int, matrix string) {
	kdx := 0
	newmatrix := matrix

	chars := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	for idx := 0; idx+x < rows; idx++ {
		for cdx := 0; cdx < cols; cdx++ {
			for jdx := 0; jdx+y < cols; jdx += 1 + cdx*idx {
				kdx = getKdx(idx+x, jdx+y, cols)
				newmatrix = newmatrix[:kdx] + chars[cdx] + newmatrix[kdx+1:]
			}
		}
	}
	kdx = getKdx(x, y, cols)
	newmatrix = newmatrix[:kdx] + "X" + newmatrix[kdx+1:]
	dump(rows, cols, &newmatrix)
}

func drawSlopes(x int, y int, rows int, cols int, matrix string) {
	kdx := 0
	newmatrix := matrix

	chars := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	for cdx := 0; cdx < cols; cdx++ {
		for idx := 0; idx+x < rows; idx++ {
			for jdx := 0; jdx+y < cols; jdx += idx + cdx*idx {
				kdx = getKdx(idx+x, jdx+y, cols)
				newmatrix = newmatrix[:kdx] + chars[cdx] + newmatrix[kdx+1:]
			}
		}
	}
	kdx = getKdx(x, y, cols)
	newmatrix = newmatrix[:kdx] + "X" + newmatrix[kdx+1:]
	dump(rows, cols, &newmatrix)
}

func parseTwo(rows int, cols int, matrix string) {
	for idx := 0; idx < rows; idx++ {
		for jdx := 0; jdx < cols; jdx++ {
			drawSlopes(idx, jdx, rows, cols, matrix)
			break
		}
		break
	}
}

// -----------------------------------------------------------------

func parse(lines []string) {
	// fmt.Println(lines)
	matrix := ""

	rows := len(lines)
	for idx := 0; idx < rows; idx++ {
		matrix += strings.ReplaceAll(lines[idx], "\n", "")
	}
	cols := len(matrix) / rows

	parseOne(rows, cols, matrix)
	parseTwo(rows, cols, matrix)

	fmt.Println("rows ", rows, " - cols ", cols)
	// dump(rows, cols, &matrix)
}

// -----------------------------------------------------------------

func main() {
	lines := helper.FileAsStringList()
	parse(lines)
}
