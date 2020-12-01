package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)

	arr := []int64{}
	for s.Scan() {
		if i, err := strconv.ParseInt(s.Text(), 10, 0); err == nil {
			arr = append(arr, i)
			// fmt.Println(i)
		}
	}

	// fmt.Printf("arr=%v\n", arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			// fmt.Printf("i=%d j=%d i+j=%d\n", arr[i], arr[j], arr[i]+arr[j])
			if (arr[i] + arr[j]) == 2020 {
				fmt.Printf("i=%d j=%d i+j=%d i*j=%d\n", arr[i], arr[j], arr[i]+arr[j], arr[i]*arr[j])
			}
		}
		// fmt.Println(arr[i])
	}

	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				// fmt.Printf("i=%d j=%d i+j=%d\n", arr[i], arr[j], arr[i]+arr[j])
				if (arr[i] + arr[j] + arr[k]) == 2020 {
					fmt.Printf("i=%d j=%d k=%d i+j+k=%d i*j*k=%d\n", arr[i], arr[j], arr[k], arr[i]+arr[j]+arr[k], arr[i]*arr[j]*arr[k])
				}
			}
		}
		// fmt.Println(arr[i])
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
