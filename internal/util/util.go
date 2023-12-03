package util

import (
	"flag"
	"fmt"
)

func Init(one func(i string), two func(i string)) {
	var input string
	var part int
	flag.StringVar(&input, "input", "./test_input.txt", "name of input file")
	flag.IntVar(&part, "part", 0, "part of day")
	flag.Parse()

	fmt.Printf("input file: %v\n", input)

	switch part {
	case 1:
		one(input)
	case 2:
		two(input)
	default:
		one(input)
	}
}
