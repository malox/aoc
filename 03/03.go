package main

import (
	"fmt"
	helper "helper"
)

func main() {

	arr := helper.FileAsStringList()
	//fmt.Printf("arr=%v\n", arr)
	
	rows := len(arr)
	cols := len(arr[0])
	fmt.Printf("rows %d - cols %d\n", rows, cols)

	tree := []int{}

	incr_i := []int{1,3,5,7,1}
	incr_j := []int{1,1,1,1,2}

	for r := 0; r < len(incr_i) ; r++ {
		i := 0
		j := 0
		curr := 0
		for j < rows {
			i += incr_i[r]
			if i >= cols {
				i -= cols
			}
			j += incr_j[r]
		
			//fmt.Printf("j %d - i %d\n", j, i)
			if j < rows && arr[j][i] == '#' {
				curr++
			}
		}
		tree = append(tree, curr)
	}

	tot := 1
	for r := 0; r < len(incr_i) ; r++ {
		fmt.Printf("right %d down %d - tree %d\n", incr_i[r], incr_j[r], tree[r])
		tot *= tree[r]
	}
	fmt.Printf("tot %d \n", tot)
}