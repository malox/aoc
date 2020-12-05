package main

import (
	"fmt"
	helper "helper"
	"sort"
)

func main() {

	arr := helper.FileAsStringList()
	//fmt.Printf("arr=%v\n", arr)

	bpmax := 0

	allseats := []int{}

	for r := 0; r < len(arr); r++ {
		bp := arr[r]
		//fmt.Printf("\n\nbp=%s\n", bp)
		rowmin := 0
		rowmax := 127
		for i := 0; i < 7; i++ {
			if bp[i] == 'F' {
				rowmax -= (rowmax + 1 - rowmin) / 2
			} else {
				rowmin += (rowmax + 1 - rowmin) / 2
			}
			//fmt.Printf(" - char=%c rowmin=%d rowmax=%d\n", bp[i], rowmin, rowmax)
		}

		seatmin := 0
		seatmax := 7
		for i := 7; i < 10; i++ {
			if bp[i] == 'L' {
				seatmax -= (seatmax + 1 - seatmin) / 2
			} else {
				seatmin += (seatmax + 1 - seatmin) / 2
			}
			//fmt.Printf(" - char=%c seatmin=%d seatmax=%d\n", bp[i], seatmin, seatmax)
		}

		seatID := (rowmax * 8) + seatmax
		allseats = append(allseats, seatID)
		// fmt.Printf(" - row=%d col=%d seatID=%d\n", rowmax, seatmax, seatID)

		if seatID > bpmax {
			bpmax = seatID
		}
	}

	myseat := 0
	sort.Ints(allseats)
	for i := 1; i < len(allseats) && myseat == 0; i++ {
		if (allseats[i-1] + 1) != allseats[i] {
			myseat = allseats[i-1] + 1
		}
	}

	fmt.Printf("bpmax=%d myseat=%d\n\n", bpmax, myseat)
}
