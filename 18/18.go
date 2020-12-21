package main

import (
	"fmt"
	"helper"
	"strings"
)

type expr []interface{}

// -----------------------------------------------------------------

func calc(one, two int64, ops string) int64 {
	if ops == "*" {
		return one * two
	} else if ops == "+" {
		return one + two
	}
	return 0
}

// -----------------------------------------------------------------

func eval(e expr) int64 {
	if len(e) == 0 {
		return int64(0)
	}
	var tot int64
	exp := e[0]
	switch exp.(type) {
	case string:
		val, _ := helper.Atoi64(exp.(string))
		tot = val
	case expr:
		tot = eval(exp.(expr))
	}

	var ops string
	for idx := 1; idx < len(e); idx++ {
		exp = e[idx]
		switch exp.(type) {
		case string:
			val, err := helper.Atoi64(exp.(string))
			if err == nil {
				tot = calc(tot, val, ops)
			} else {
				ops = exp.(string)
			}
		case expr:
			tot = calc(tot, eval(exp.(expr)), ops)
		}
	}

	return tot
}

// -----------------------------------------------------------------

func evaltwo(e expr) int64 {
	if len(e) == 0 {
		return int64(0)
	}
	var tot int64
	exp := (e)[0]
	switch exp.(type) {
	case string:
		val, _ := helper.Atoi64(exp.(string))
		tot = val
	case expr:
		tot = evaltwo(exp.(expr))
	}

	e = e[1:]
	var ops string
	for len(e) > 0 {
		ops = e[0].(string)
		exp = e[1]
		if ops == "+" {
			switch exp.(type) {
			case string:
				val, _ := helper.Atoi64(exp.(string))
				tot = calc(tot, val, ops)
			case expr:
				tot = calc(tot, evaltwo(exp.(expr)), ops)
			}
			e = e[2:]
		} else if ops == "*" {
			tmpval := evaltwo(e[1:])
			tot = calc(tot, tmpval, ops)
			e = e[len(e):]
		} else {
			fmt.Println("evaltwo: unexpected val for ops ", ops)
		}
	}

	return tot
}

// -----------------------------------------------------------------

func parse(lines [][]string) {
	input := []expr{}

	for _, line := range lines {

		// fmt.Println("parsing line:", line)
		stack := []expr{}
		stackidx := 0
		currexpr := expr{}
		for _, val := range line {
			if val == "(" {
				stackidx++
				stack = append(stack, currexpr)
				currexpr = expr{}
			} else if val == ")" {
				stackidx--
				tmpexpr := currexpr
				currexpr = stack[stackidx]
				currexpr = append(currexpr, tmpexpr)
				stack = stack[:stackidx]
			} else {
				currexpr = append(currexpr, val)
			}
		}
		input = append(input, currexpr)
	}

	fmt.Printf("%s\n\n", strings.Repeat("-", 23))
	var totsum int64
	for _, expr := range input {
		currtot := eval(expr)
		// fmt.Println("currtot : ", currtot, " - ", expr)
		totsum += currtot
	}
	fmt.Printf("\ntotsum : %d \n\n", totsum)

	fmt.Printf("%s\n\n", strings.Repeat("-", 23))
	totsum = 0
	for _, expr := range input {
		currtot := evaltwo(expr)
		// fmt.Println("currtot2 : ", currtot, " - ", expr)
		totsum += currtot
	}
	fmt.Printf("\ntotsum2 : %d \n\n", totsum)
}

// -----------------------------------------------------------------

func main() {
	lines := helper.FileAsStringList()
	var newlines [][]string
	for _, line := range lines {
		patched := strings.ReplaceAll(strings.ReplaceAll(line, "(", "( "), ")", " )")
		newlines = append(newlines, strings.Split(patched, " "))
	}
	parse(newlines)
}
