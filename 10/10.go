package main

import (
	"fmt"
	"helper"
	"sort"
)

// -----------------------------------------------------------------

type adapter struct {
	val   int
	nexts map[int]int
}

func getAdapter(val int, adapters *map[int]adapter) adapter {
	local, exists := (*adapters)[val]
	if !exists {
		local = adapter{}
		local.val = val
		local.nexts = map[int]int{}
		(*adapters)[val] = local
	}
	return local
}

func dumpAdapters(adapters *map[int]adapter) {
	for k, v := range *adapters {
		fmt.Println(" - id ", k, "- [", v.nexts, "]")
	}
}

func sortAdapters(adapters *map[int]adapter) []int {
	arr := []int{}
	for k := range *adapters {
		arr = append(arr, k)
	}
	sort.IntSlice(arr).Sort()
	return arr
}

// -----------------------------------------------------------------

// part 2.1

func calcAdapter(id int, adapters *map[int]adapter, count *map[int]int) {
	children := len((*adapters)[id].nexts)
	if children == 0 {
		(*count)[id] = 1
	} else {
		sum := 0
		for _, child := range (*adapters)[id].nexts {
			sum += (*count)[child]
		}
		(*count)[id] = sum
	}
}

func findAll(idx int, values *[]int, adapters *map[int]adapter, count *map[int]int) {

	for idx := len(*values) - 1; idx >= 0; idx-- {
		val := (*values)[idx]
		calcAdapter(val, adapters, count)
	}
}

// part 2.2
func paths(start int, stop int, adapters *map[int]adapter, count *map[int]int) int {

	if start == stop {
		return 1
	}

	if val, exists := (*count)[start]; exists {
		return val
	}

	tot := 1
	children := (*adapters)[start].nexts
	if len(children) > 0 {
		tot = 0
		offset := []int{1, 2, 3}
		for idx := range offset {
			nextid, exists := children[offset[idx]]
			if exists {
				next := (*adapters)[nextid]
				tot += paths(next.val, stop, adapters, count)
			}
		}
	}
	(*count)[start] = tot
	return tot
}

// -----------------------------------------------------------------

// part one

func findDiff(start int, stop int, offset []int, adapters *map[int]adapter) {
	diff := map[int]int{1: 0, 2: 0, 3: 0}
	iter := (*adapters)[start]
	end := (*adapters)[stop]
	for iter.val != end.val {
		for idx := range offset {
			next, exists := iter.nexts[offset[idx]]
			if exists {
				diff[offset[idx]]++
				iter = (*adapters)[next]
				break
			}
		}
	}
	fmt.Println(diff, " diff ", diff[1]*diff[3])
}

// -----------------------------------------------------------------

func parse(intlist []int) {
	// fmt.Println(intlist)

	adapters := map[int]adapter{}
	getAdapter(0, &adapters)

	max := 0
	for idx := 0; idx < len(intlist); idx++ {
		curr := getAdapter(intlist[idx], &adapters)
		if curr.val > max {
			max = intlist[idx]
		}
	}

	max += 3
	end := getAdapter(max, &adapters)

	offset := []int{1, 2, 3}
	for _, iter := range adapters {
		for idx := range offset {
			nextpos := iter.val + offset[idx]
			_, exists := adapters[nextpos]
			if exists {
				iter.nexts[offset[idx]] = nextpos
			}
		}
	}

	// dumpAdapters(&adapters)

	findDiff(0, end.val, offset, &adapters)

	sorted := sortAdapters(&adapters)
	// fmt.Println(sorted)

	count := map[int]int{}
	findAll(len(sorted), &sorted, &adapters, &count)
	fmt.Println("tot2 ", count[0])

	countbis := map[int]int{}
	fmt.Println("tot3 ", paths(0, max, &adapters, &countbis))
}

func main() {
	intlist := helper.FileAsIntList()
	parse(intlist)
}
