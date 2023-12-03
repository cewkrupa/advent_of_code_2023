package util

import (
	"flag"
	"fmt"
)

func Init(one func(fileName string), two func(fileName string)) {
	var fileName string
	var part int
	flag.StringVar(&fileName, "fileName", "./test_input.txt", "name of input file")
	flag.IntVar(&part, "part", 0, "part of day")
	flag.Parse()

	fmt.Printf("input file: %v\n", fileName)

	switch part {
	case 1:
		one(fileName)
	case 2:
		two(fileName)
	default:
		one(fileName)
	}
}
