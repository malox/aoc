package main

import (
	"fmt"
	"helper"
	"strconv"
)

func updateDir(dir *string, degree int) {
	for steps := degree / 90; steps != 0; steps-- {
		if *dir == "N" {
			*dir = "W"
		} else if *dir == "S" {
			*dir = "E"
		} else if *dir == "E" {
			*dir = "N"
		} else if *dir == "W" {
			*dir = "S"
		}
	}
}

func myAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func move(x *int, y *int, offset int, dir byte) {
	if dir == 'N' {
		*y += offset
	} else if dir == 'S' {
		*y -= offset
	} else if dir == 'E' {
		*x += offset
	} else if dir == 'W' {
		*x -= offset
	}
}

func parse(commands []string) {
	x := 0
	y := 0
	dir := "E"
	for idx := 0; idx < len(commands); idx++ {
		cmd := commands[idx][0]
		offset, _ := strconv.Atoi(commands[idx][1:])
		if cmd == 'F' {
			cmd = dir[0]
		}

		if cmd == 'L' {
			updateDir(&dir, offset)
		} else if cmd == 'R' {
			updateDir(&dir, 360-offset)
		} else {
			move(&x, &y, offset, cmd)
		}
	}
	fmt.Println("one ", myAbs(x)+myAbs(y))
}

func updateWpointPos(x *int, y *int, degree int) {

	for steps := degree / 90; steps != 0; steps-- {
		tmpx := *x
		*x = -(*y)
		*y = tmpx
	}
}

func parseTwo(commands []string) {
	shipx := 0
	shipy := 0
	wpointx := 10
	wpointy := 1

	for idx := 0; idx < len(commands); idx++ {
		cmd := commands[idx][0]
		offset, _ := strconv.Atoi(commands[idx][1:])
		if cmd == 'F' {
			shipx += offset * wpointx
			shipy += offset * wpointy
		} else if cmd == 'L' {
			updateWpointPos(&wpointx, &wpointy, offset)
		} else if cmd == 'R' {
			updateWpointPos(&wpointx, &wpointy, 360-offset)
		} else {
			move(&wpointx, &wpointy, offset, cmd)
		}
		// fmt.Printf(" ship %d %d - wp %d %d \n ", shipx, shipy, wpointx, wpointy)
	}
	fmt.Println("two ", myAbs(shipx)+myAbs(shipy))
}

func main() {
	lines := helper.FileAsStringList()
	parse(lines)
	parseTwo(lines)
}
