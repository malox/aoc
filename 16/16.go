package main

import (
	"fmt"
	"helper"
	"strconv"
	"strings"
)

// -----------------------------------------------------------------

func ticketToInts(ticket string) []int {
	var ints []int
	vals := strings.Split(ticket, ",")
	for idx := range vals {
		num, _ := strconv.Atoi(vals[idx])
		ints = append(ints, num)
	}
	return ints
}

func validField(val int, rule *[]int) bool {
	for jdx := 0; jdx < 3; jdx += 2 {
		if val >= (*rule)[jdx] && val <= (*rule)[jdx+1] {
			return true
		}
	}
	return false
}

func validTicket(val int, rules *map[string][]int) bool {
	for _, rule := range *rules {
		if validField(val, &rule) {
			return true
		}
	}
	return false
}

func contains(arr []int, x int) bool {
	for _, n := range arr {
		if x == n {
			return true
		}
	}
	return false
}

// -----------------------------------------------------------------

func parse(lines []string) {
	// fmt.Println(lines)

	parsemine := false
	var myticket []int
	parsenearby := false
	var nearbytickets [][]int
	rules := map[string][]int{}
	fields := map[string][]int{}

	invalidrange := 0
	for _, line := range lines {
		if parsenearby {
			ints := ticketToInts(line)
			valid := true
			for _, val := range ints {
				if !validTicket(val, &rules) {
					invalidrange += val
					valid = false
				}
			}
			if valid {
				nearbytickets = append(nearbytickets, ints)
			}
			continue
		}
		if strings.Contains(line, "nearby tickets:") {
			parsenearby = true
			continue
		}

		if parsemine {
			myticket = ticketToInts(line)
			parsemine = false
			continue
		}
		if strings.Contains(line, "your ticket:") {
			parsemine = true
			continue
		}

		var idx1, idx2, idx3, idx4 int
		flddesc := strings.Split(line, ":")
		tag := flddesc[0]
		fmt.Sscanf(flddesc[1], " %d-%d or %d-%d", &idx1, &idx2, &idx3, &idx4)
		// fmt.Printf("%s -> %d-%d | %d-%d\n", tag, idx1, idx2, idx3, idx4)
		rules[tag] = []int{idx1, idx2, idx3, idx4}
	}

	// fmt.Println("rules ", rules)
	// fmt.Println("myticket ", myticket)
	// fmt.Println("nearbytickets ", nearbytickets)
	fmt.Println("invalidrange ", invalidrange)

	for key := range rules {
		fields[key] = []int{}
	}

	for key, rule := range rules {
		for jdx := range myticket {
			valid := true
			for _, ticket := range nearbytickets {
				val := ticket[jdx]
				if !validField(val, &rule) {
					valid = false
					break
				}
			}
			if valid {
				fields[key] = append(fields[key], jdx)
			}
		}
	}

	// fmt.Println("fields ", fields)

	for true {
		singlerule := []int{}
		for _, rule := range fields {
			if len(rule) == 1 {
				singlerule = append(singlerule, rule[0])
			}
		}

		if len(singlerule) == len(fields) {
			break
		}

		for key, rule := range fields {
			if len(rule) == 1 {
				continue
			}
			newrule := []int{}
			for _, val := range rule {
				if !contains(singlerule, val) {
					newrule = append(newrule, val)
				}
			}
			fields[key] = newrule
		}
	}

	depart := 1
	for key, rule := range fields {
		// sort.Ints(rule)
		// fmt.Println("key ", key, " - valid ", rule)
		if strings.Contains(key, "departure") {
			depart *= myticket[rule[0]]
		}
	}
	fmt.Println("departure ", depart)
}

// -----------------------------------------------------------------

func main() {
	lines := helper.FileAsStringList()
	parse(lines)
}
