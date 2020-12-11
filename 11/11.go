package main

import (
	"fmt"
	"helper"
	"strings"
)

// -----------------------------------------------------------------

func dump(rows int, cols int, matrix *string) {
	fmt.Println(strings.Repeat("-", cols))
	for idx := 0; idx < rows; idx++ {
		for jdx := 0; jdx < cols; jdx++ {
			kdx := idx*cols + jdx
			fmt.Printf("%c", (*matrix)[kdx])
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
			kdx := idx*cols + jdx
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
				kdx := idx*cols + jdx
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

func parseOneBis(rows int, cols int, matrix string) {

	oldmatrix := ""
	for matrix != oldmatrix {
		oldmatrix = matrix
		for idx := 0; idx < rows; idx++ {
			for jdx := 0; jdx < cols; jdx++ {
				kdx := idx*cols + jdx
				seat := matrix[kdx]
				busy := occupy(idx, jdx, rows, cols, &matrix)
				if seat == 'L' && busy == 0 {
					matrix = matrix[:kdx] + "#" + matrix[kdx+1:]
				} else if seat == '#' && busy >= 4 {
					matrix = matrix[:kdx] + "L" + matrix[kdx+1:]
				}
			}
		}
		dump(rows, cols, &matrix)
	}

	fmt.Println("parseOne - occupied seats ", strings.Count(matrix, "#"))

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
	// parseOneBis(rows, cols, matrix)

	fmt.Println("rows ", rows, " - cols ", cols)
	// dump(rows, cols, &matrix)
}

// -----------------------------------------------------------------

func main() {
	lines := helper.FileAsStringList()
	parse(lines)
}
