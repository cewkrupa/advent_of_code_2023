package v1

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"
)

var digitRegex = regexp.MustCompile(`(\d+)`)

const DigitPattern = `(\d+)`

func DigitRegex() *regexp.Regexp {
	return regexp.MustCompile(DigitPattern)
}

func Init(one func(inputFile *os.File), two func(inputFile *os.File)) {
	startTime := time.Now()

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
	endTime := time.Now()
	fmt.Printf("completed in %v ms\n", endTime.Sub(startTime).Milliseconds())
}
