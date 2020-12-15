package main

import (
	"fmt"
)

func getHist(hist *map[int][]int, value int) []int {
	_, exists := (*hist)[value]
	if !exists {
		(*hist)[value] = []int{}
	}
	return (*hist)[value]
}

func updateHist(hist *map[int][]int, value int, round int) {
	arr := getHist(hist, value)
	arr = append(arr, round)
	if len(arr) > 4 {
		arr = arr[len(arr)-3:]
	}
	(*hist)[value] = arr
}

func dumpHist(hist *map[int][]int) {
	for idx, arr := range *hist {
		fmt.Println(" - dump ", idx, " - hist ", arr)
	}
}

func parse(arr []int, stop int) {
	// fmt.Println(arr)

	hist := map[int][]int{}
	round := 1
	var spoken int

	for idx := range arr {
		spoken := arr[idx]
		updateHist(&hist, spoken, round)
		// fmt.Println("init - round ", round, " - spoken ", spoken)
		round++
	}

	// dumpHist(&hist)

	for true {
		last := spoken
		spoken = 0
		_, exists := hist[last]
		if exists {
			arr := getHist(&hist, last)
			past := len(arr)
			if past > 1 {
				spoken = (arr)[past-1] - (arr)[past-2]
			}
		}
		updateHist(&hist, spoken, round)
		// fmt.Println("game - round ", round, " - spoken ", spoken, " -- exists", exists)
		if round == stop {
			break
		}
		round++
	}

	fmt.Println("end - round ", round, " - spoken ", spoken)
}

func main() {

	input := [][]int{
		// {0, 3, 6},

		// {1, 3, 2},
		// {2, 1, 3},
		// {1, 2, 3},
		// {2, 3, 1},
		// {3, 2, 1},
		// {3, 1, 2},

		{2, 20, 0, 4, 1, 17},
	}

	for idx := range input {
		parse(input[idx], 2020)
	}
	for idx := range input {
		parse(input[idx], 30000000)
	}

}
