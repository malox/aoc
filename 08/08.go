package main

import (
	"fmt"
	helper "helper"
	"strconv"
	"strings"
)

func updateIncr(index *int, incr int, sign byte) {
	if sign == '-' {
		*index -= incr
	} else {
		*index += incr
	}
}

func parseInstr(instrlist []string) {

	acc := 0
	idx := 0
	dowork := true

	parsed := []int{}

	for dowork {

		curr := strings.Split(instrlist[idx], " ")

		instr := curr[0]
		sign := curr[1][0]
		incr, _ := strconv.Atoi(curr[1][1:])

		if instr == "nop" {
			idx++
		} else if instr == "acc" {
			updateIncr(&acc, incr, sign)
			idx++
		} else if instr == "jmp" {
			updateIncr(&idx, incr, sign)
		} else {
			fmt.Println("Found an invalid instructions at idx ", idx, " -> ", curr)
		}

		for i := 0; i < len(parsed); i++ {
			if parsed[i] == idx {
				fmt.Println("Infinite loop detected at idx ", idx, " - acc ", acc)
				dowork = false
			}
		}
		parsed = append(parsed, idx)
		if idx >= len(instrlist) {
			fmt.Println("Index out of range - exit ", idx, " - acc ", acc)
			dowork = false
		}
	}
}

func patchAndParse(idx int, instrlist []string) {
	if strings.Contains(instrlist[idx], "nop") {
		instrlist[idx] = strings.Replace(instrlist[idx], "nop", "jmp", 1)
		parseInstr(instrlist)
		instrlist[idx] = strings.Replace(instrlist[idx], "jmp", "nop", 1)
	} else if strings.Contains(instrlist[idx], "jmp") {
		instrlist[idx] = strings.Replace(instrlist[idx], "jmp", "nop", 1)
		parseInstr(instrlist)
		instrlist[idx] = strings.Replace(instrlist[idx], "nop", "jmp", 1)
	}
}

func main() {

	instrlist := helper.FileAsStringList()

	// part one
	parseInstr(instrlist)

	// part two
	for i := 0; i < len(instrlist); i++ {
		patchAndParse(i, instrlist)
	}

}
