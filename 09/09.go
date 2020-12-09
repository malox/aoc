package main

import (
	"fmt"
	"helper"
)

func resetIdx(index *int, base int, offset int) {
	*index = 0
	if base > offset {
		*index = base - offset
	}
}

func parse(intlist []int) {

	jdx := 0
	kdx := 0
	tmp := []int{}
	offset := 25
	if len(intlist) < 25 {
		offset = 5
	}

	for idx, num := range intlist {
		tmpsize := len(tmp)
		for resetIdx(&jdx, idx, offset); jdx < tmpsize; jdx++ {
			for resetIdx(&kdx, idx, offset); kdx < tmpsize; kdx++ {
				if jdx != kdx && num == intlist[jdx]+intlist[kdx] {
					tmp = append(tmp, 1)
					break
				}
			}
			if len(tmp) > idx {
				break
			}
		}
		if len(tmp) == idx {
			tmp = append(tmp, 0)
		}
	}

	// fmt.Println(intlist)
	// fmt.Println(tmp)

	invalid := 0
	for idx, num := range tmp {
		if num == 0 {
			invalid = intlist[idx]
			fmt.Printf("found 0 at idx %d corresponding to %d\n", idx, invalid)
		}
	}

	// part two
	sum := 0
	intlistsize := len(intlist)

	for idx := 0; idx < intlistsize-1; idx++ {
		for jdx = idx + 1; jdx < intlistsize; jdx++ {
			min := invalid
			max := 0
			for kdx = idx; kdx <= jdx; kdx++ {
				if intlist[kdx] > max {
					max = intlist[kdx]
				}
				if intlist[kdx] < min {
					min = intlist[kdx]
				}
				sum += intlist[kdx]
			}
			if sum == invalid {
				fmt.Printf("range idx %d jdx %d - min %d max %d - min+max %d\n", idx, jdx, min, max, min+max)
				break
			} else {
				sum = 0
			}
		}
		if sum != 0 {
			break
		}
	}
}

func main() {
	intlist := helper.FileAsIntList()
	parse(intlist)
}
