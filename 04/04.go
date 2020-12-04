package main

import (
	"fmt"
	helper "helper"
	"strings"
)

func checkTwo(passport string) bool {
	keyvals := strings.Split(passport, " ")
	for i := 0; i < len(keyvals); i++ {
		curr := strings.Split(keyvals[i], ":")
		key := curr[0]
		val := curr[1]
		switch key {
		case "byr":
			if val < "1920" || val > "2002" {
				fmt.Printf("birth year - %s - %s\n", key, val)
				return false
			}
		case "iyr":
			if val < "2010" || val > "2020" {
				fmt.Printf("issue year - %s - %s\n", key, val)
				return false
			}
		case "eyr":
			if val < "2020" || val > "2030" {
				fmt.Printf("expire year - %s - %s\n", key, val)
				return false
			}
		case "hgt":
			// a number followed by either cm or in:
			// If cm, the number must be at least 150 and at most 193.
			// If in, the number must be at least 59 and at most 76.
			hgt := val[:len(val)-2]
			if strings.HasSuffix(val, "cm") {
				if hgt < "150" || hgt > "193" {
					fmt.Printf("height cm %s - %s - %s\n", hgt, key, val)
					return false
				}
			} else if strings.HasSuffix(val, "in") {
				if hgt < "59" || hgt > "76" {
					fmt.Printf("height in - %s - %s\n", key, val)
					return false
				}
			} else {
				fmt.Printf("invalid height - %s - %s\n", key, val)
				return false
			}
		case "hcl":
			// a # followed by exactly six characters 0-9 or a-f
			if len(val) == 7 && val[0] == '#' {
				for k := 1; k < len(val); k++ {
					ch := val[k]
					valid := (ch >= '0' || ch <= '9') || (ch >= 'a' || ch <= 'f')
					if !valid {
						fmt.Printf("invalid hair color char %c - %s - %s\n", ch, key, val)
						return false
					}
				}
			} else {
				fmt.Printf("invalid hair color - %s - %s\n", key, val)
				return false
			}
		case "ecl":
			eyecolors := "amb blu brn gry grn hzl oth"
			if !strings.Contains(eyecolors, val) {
				fmt.Printf("eye color - %s - %s\n", key, val)
				return false
			}
		case "pid":
			// nine-digit number
			if len(val) == 9 {
				for k := 0; k < len(val); k++ {
					ch := val[k]
					valid := (ch >= '0' || ch <= '9')
					if !valid {
						fmt.Printf("invalid passport char %c - %s - %s\n", ch, key, val)
						return false
					}
				}
			} else {
				fmt.Printf("invalid passport id - %s - %s\n", key, val)
				return false
			}
		case "cid":
			//fmt.Printf("skipped key - %s - %s\n", key, val)
		default:
			fmt.Printf("unkown keyval - %s - %s - for passport %s\n", key, val, passport)
		}
	}
	return len(passport) > 0
}

func main() {

	arr := helper.FileAsPassportList()
	//fmt.Printf("arr=%v\n", arr)

	totOne := 0
	totTwo := 0
	keys := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:" /*, "cid"*/}

	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(keys); j++ {
			if strings.Contains(arr[i], keys[j]) {
				count++
			}
		}
		if count == len(keys) {
			//mt.Printf("valid = %s\n", arr[i])
			totOne++
			if checkTwo(arr[i]) {
				totTwo++
			} else {
				fmt.Printf("invalid = %s\n\n", arr[i])
			}

		} //else {
		//	fmt.Printf("wrong = %s\n", arr[i])
		//}
	}
	fmt.Printf("totOne=%d - totTwo=%d\n", totOne, totTwo)
}
