package main

import (
	"fmt"
	helper "helper"
	"strings"
)

func addto(str *string, arr *[]string) {
	if len(*str) > 0 {
		// fmt.Printf("append %s\n", *str)
		*arr = append(*arr, *str)
		*str = ""
	}
}

func main() {

	fs := helper.OpenFile()

	str := ""
	arr := []string{}

	partwo := []string{}

	for fs.Scan() {

		tmp := fs.Text()
		partwo = append(partwo, tmp)
		if len(tmp) > 0 {
			for i := 0; i < len(tmp); i++ {
				if !strings.Contains(str, string(tmp[i])) {
					str += string(tmp[i])
				}
			}
		} else {
			addto(&str, &arr)
		}
	}

	addto(&str, &arr)

	count := 0
	for i := 0; i < len(arr); i++ {
		count += len(arr[i])
	}
	fmt.Printf("count %d\n\n", count)

	str = ""
	first := true
	arrtwo := []string{}

	for k := 0; k < len(partwo); k++ {

		tmp := partwo[k]
		// fmt.Printf("\n extracted from vector [%s]\n", tmp)

		if len(tmp) > 0 {
			if first {
				str = tmp
				first = false
				// fmt.Printf("first %s\n", str)
			} else if len(str) > 0 {
				tmpStr := ""
				for i := 0; i < len(tmp); i++ {
					if strings.Contains(str, string(tmp[i])) {
						tmpStr += string(tmp[i])
					}
				}
				str = tmpStr
				// fmt.Printf("next %s - str %s\n", tmp, str)
			}
		} else {
			addto(&str, &arrtwo)
			first = true
		}
	}

	addto(&str, &arrtwo)

	count = 0
	for i := 0; i < len(arrtwo); i++ {
		count += len(arrtwo[i])
	}
	fmt.Printf("count_two %d\n\n", count)
}
