package helper

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func openFile() *bufio.Scanner {

	fptr := flag.String("fp", "test", "file path to read from")
	flag.Parse()

	fmt.Println("opening file : " + *fptr + "\n")

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	// defer func() {
	// 	if err = f.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }() // not closing the file
	return bufio.NewScanner(f)
}

func FileAsIntList() []int64 {

	fs := openFile()

	arr := []int64{}
	for fs.Scan() {
		if i, err := strconv.ParseInt(fs.Text(), 10, 0); err == nil {
			arr = append(arr, i)
			// fmt.Println(i)
		}
	}

	err := fs.Err()
	if err != nil {
		log.Fatal(err)
	}

	return arr
}

func FileAsStringList() []string {

	fs := openFile()

	arr := []string{}
	for fs.Scan() {
		if len(fs.Text()) > 0 {
			arr = append(arr, fs.Text())
			// 	fmt.Println("adding : " + fs.Text())
			// } else {
			// 	fmt.Println("skip empty line" )
		}
	}

	err := fs.Err()
	if err != nil {
		log.Fatal(err)
	}

	return arr
}

func addPassport(arr *[]string, passport *string) {
	//fmt.Println("Adding passport : " + *passport + "\n")
	*arr = append(*arr, *passport)
	*passport = ""
}

func FileAsPassportList() []string {

	fs := openFile()

	arr := []string{}

	tmpPassport := ""

	for fs.Scan() {
		if len(fs.Text()) > 0 {
			if len(tmpPassport) > 0 {
				tmpPassport += " "
			}
			tmpPassport += fs.Text()
			//fmt.Println(" - adding passport chunk : " + fs.Text())
		} else {
			addPassport(&arr, &tmpPassport)
		}
	}
	if len(tmpPassport) > 0 {
		addPassport(&arr, &tmpPassport)
	}

	err := fs.Err()
	if err != nil {
		log.Fatal(err)
	}

	return arr
}
