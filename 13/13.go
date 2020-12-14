package main

import (
	"fmt"
	"helper"
	"strconv"
	"strings"
)

// -----------------------------------------------------------------

// -----------------------------------------------------------------

func parse(lines []string) {
	// fmt.Println(lines)
	time, _ := strconv.Atoi(lines[0])
	buses := strings.Split(lines[1], ",")

	minwait := time
	minwaitbus := 0
	for idx := 0; idx < len(buses); idx++ {
		bus, err := strconv.Atoi(buses[idx])
		if err == nil {
			previous := time / bus
			next := bus * (previous + 1)
			wait := next - time
			if wait < minwait {
				minwait = wait
				minwaitbus = bus
			}
		}
	}

	fmt.Println("time ", time, " tot ", minwait*minwaitbus)
	// fmt.Println("table ", table)
}

func parseTwo(lines []string) {
	// fmt.Println(lines)
	for idx := range lines {
		buses := strings.Split(lines[idx], ",")
		size := len(buses)
		if size < 2 {
			continue
		}
		fmt.Println(buses)

		var time int64
		factor, _ := helper.Atoi64(buses[0])
		jdx := 1

		for jdx < size {
			curr, err := helper.Atoi64(buses[jdx])
			if err != nil {
				jdx++
				continue
			}
			time += factor
			if (time+int64(jdx))%curr == 0 {
				factor *= curr
				jdx++
			}
		}

		fmt.Println("found min at ", time, " for buses ", buses)
	}
}

// -----------------------------------------------------------------

func main() {
	lines := helper.FileAsStringList()
	parse(lines[:2])
	parseTwo(lines)
}
