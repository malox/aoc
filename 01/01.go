package main

import (
	"fmt"
	helper "helper"
)

func main() {

	arr := helper.FileAsIntList()

	// fmt.Printf("arr=%v\n", arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			// fmt.Printf("i=%d j=%d i+j=%d\n", arr[i], arr[j], arr[i]+arr[j])
			if (arr[i] + arr[j]) == 2020 {
				fmt.Printf("i=%d j=%d i+j=%d i*j=%d\n", arr[i], arr[j], arr[i]+arr[j], arr[i]*arr[j])
			}
		}
	}

	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if (arr[i] + arr[j] + arr[k]) == 2020 {
					fmt.Printf("i=%d j=%d k=%d i+j+k=%d i*j*k=%d\n", arr[i], arr[j], arr[k], arr[i]+arr[j]+arr[k], arr[i]*arr[j]*arr[k])
				}
			}
		}
	}

}
