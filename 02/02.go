package main

import (
	"fmt"
	"strings"
	"strconv"
	helper "helper"
)

func main() {

	arr := helper.FileAsStringList()

	//fmt.Printf("arr=%v\n", arr)
	
	tot := 0
	newtot := 0

	for i := 0; i < len(arr); i++ {

		curr := strings.Split(arr[i], ": ")
		//fmt.Printf("%q\n", curr)

		pwd := curr[1] 
	
		key := strings.Split(curr[0], " ")

		ch := key[1]

		policy := strings.Split(key[0], "-")

		min, _ := strconv.Atoi(policy[0])
		max, _ := strconv.Atoi(policy[1])


		count := strings.Count(pwd, ch)

		if count >= min && count <= max {
			tot++
		}

		if string(pwd[min-1]) == ch && string(pwd[max-1]) != ch {
			newtot++
		}
		if string(pwd[min-1]) != ch && string(pwd[max-1]) == ch {
			newtot++
		}

		//fmt.Printf("Found %s %d times within %s - tot %d - newtot %d \n", ch, count, pwd, tot, newtot)
	}

	fmt.Printf("Total match %d %d\n", tot, newtot)
	
}