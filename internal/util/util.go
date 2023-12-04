package util

import (
	"flag"
	"fmt"
	"os"
)

func Init(one func(inputFile *os.File), two func(inputFile *os.File)) {
	var fileName string
	var part int
	flag.StringVar(&fileName, "fileName", "./test_input.txt", "name of input file")
	flag.IntVar(&part, "part", 0, "part of day")
	flag.Parse()

	fmt.Printf("input file: %v\n", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	switch part {
	case 1:
		fmt.Println("one")
		one(file)
	case 2:
		fmt.Println("two")
		two(file)
	default:
		one(file)
	}
}
