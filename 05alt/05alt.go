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

		rownum := 0
		row := 64
		for i := 0; i < 7; i++ {
			if bp[i] == 'B' {
				rownum += row
			}
			row /= 2
			//fmt.Printf(" - char=%c rownum=%d row=%d\n", bp[i], rownum, row)
		}

		seatnum := 0
		seat := 4
		for i := 7; i < 10; i++ {
			if bp[i] == 'R' {
				seatnum += seat
			}
			seat /= 2
			//fmt.Printf(" - char=%c seatnum=%d seat=%d\n", bp[i], seatnum, seat)
		}

		seatID := (rownum * 8) + seatnum
		allseats = append(allseats, seatID)
		// fmt.Printf(" - row=%d col=%d seatID=%d\n", rownum, seatnum, seatID)

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
