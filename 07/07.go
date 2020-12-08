package main

import (
	"fmt"
	helper "helper"
	"strconv"
	"strings"
)

type bag struct {
	color string

	ChildBags  map[string]bagHolder
	ParentBags map[string]bagHolder
}

type bagHolder struct {
	bag   *bag
	count int
}

func getBag(color string, bags *map[string]bag) bag {

	localbag, exists := (*bags)[color]
	if !exists {
		// fmt.Println("Creating bag ", color)
		localbag = bag{}
		localbag.color = color
		localbag.ChildBags = map[string]bagHolder{}
		localbag.ParentBags = map[string]bagHolder{}
		(*bags)[color] = localbag
	}

	return localbag
}

func getBagHolder(bag *bag, count int) bagHolder {
	localBagHold := bagHolder{}
	localBagHold.bag = bag
	localBagHold.count = count
	return localBagHold
}

func doSearch(bag *bag, tot *int, totStr *string, offset string) {

	parents := (*bag).ParentBags
	// fmt.Println(offset, "searching ", (*bag).color, " - parents ", len(parents))

	if len(parents) > 0 {
		for _, v := range parents {
			doSearch(v.bag, tot, totStr, offset+"  ")
		}
	}
	if !strings.Contains(*totStr, (*bag).color) {
		*totStr += " - " + (*bag).color
		*tot++
	}
}

func doSearchBis(bag *bag, tot *int, offset string, multiplier int) {

	children := (*bag).ChildBags

	*tot += 1 * multiplier

	for _, v := range children {
		// fmt.Println(offset, "searching ", (*bag).color, " - children ", v.count, " ", k, " - curr ", *tot, " - multiplier ", multiplier)
		doSearchBis(v.bag, tot, offset+"  ", multiplier*v.count)
	}
}

func main() {

	arr := helper.FileAsStringList()
	bags := map[string]bag{}

	for k := 0; k < len(arr); k++ {
		// fmt.Println(k, " - ", arr[k])
		tmp := strings.Split(arr[k], "contain ")
		bagone := strings.Split(tmp[0], " bag")[0]

		bagit := getBag(bagone, &bags)

		tmpbags := strings.Split(tmp[1], ", ")
		for i := 0; i < len(tmpbags); i++ {
			if strings.Contains(tmpbags[i], "no other bags") {
				continue
			}
			curr := strings.Split(tmpbags[i], " ")
			count, _ := strconv.Atoi(curr[0])
			currbag := curr[1] + " " + curr[2]
			currbagit := getBag(currbag, &bags)

			currbagit.ParentBags[bagone] = getBagHolder(&bagit, 0)
			bagit.ChildBags[currbag] = getBagHolder(&currbagit, count)

		}
	}

	tot := 0
	totStr := "shiny gold"
	shinybag, exists := bags[totStr]

	if !exists {
		fmt.Println("Could not find shiny gold")
	} else {
		doSearch(&shinybag, &tot, &totStr, "  ")
	}

	fmt.Println("tot one :", tot /*, " - totStr ", totStr*/)

	if exists {
		tot = -1
		doSearchBis(&shinybag, &tot, "  ", 1)
	}

	fmt.Println("tot two:", tot)
}
