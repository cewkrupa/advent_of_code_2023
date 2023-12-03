package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)

	r, err := regexp.Compile("(\\d)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)")
	check(err)

	var runningTotal int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		digits := getMatchesFromLine(r, line, []string{})

		first, err := parseDigit(digits[0])
		check(err)

		var last int
		if len(digits) < 2 {
			last = first
		} else {
			last, err = parseDigit(digits[len(digits)-1])
			check(err)
		}

		firstAndLast := fmt.Sprintf("%v%v", first, last)
		calibrationVal, err := strconv.Atoi(firstAndLast)
		check(err)
		fmt.Println(calibrationVal)
		runningTotal += calibrationVal

	}
	fmt.Printf("%v", runningTotal)
}

func getMatchesFromLine(r *regexp.Regexp, line string, matches []string) []string {
	if len(line) <= 0 {
		return matches
	}

	lb := []byte(line)
	loc := r.FindIndex(lb)

	if loc == nil {
		// there's no more matches in the line, so return whatever we've got
		return matches
	}
	// see https://pkg.go.dev/regexp#Regexp.FindIndex
	match := lb[loc[0]:loc[1]]

	newMatches := append(matches, string(match))

	// get the matches of the line starting at the index of the first match + 1
	return getMatchesFromLine(r, string(lb[(loc[0]+1):]), newMatches)
}

func parseDigit(s string) (int, error) {

	i, err := strconv.Atoi(s)

	// we didn't get an error, so converting must have worked
	if err == nil {
		return i, nil
	}

	var result int
	switch s {
	case "one":
		result = 1
	case "two":
		result = 2
	case "three":
		result = 3
	case "four":
		result = 4
	case "five":
		result = 5
	case "six":
		result = 6
	case "seven":
		result = 7
	case "eight":
		result = 8
	case "nine":
		result = 9
	}
	if result == 0 {
		return 0, errors.New("error: unparsable string")
	}

	return result, nil
}
