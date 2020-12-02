package helper

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func FileAsIntList() []int64 {

	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	fmt.Println("opening file : " + *fptr)

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

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return arr
}



func FileAsStringList() []string {

	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	fmt.Println("opening file : " + *fptr)

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

	arr := []string{}
	for s.Scan() {
		if len(s.Text()) > 0 {
			arr = append(arr, s.Text())
		// 	fmt.Println("adding : " + s.Text())
		// } else {
		// 	fmt.Println("skip empty line" )
		}
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return arr
}

