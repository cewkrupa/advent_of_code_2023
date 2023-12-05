package main

import (
	"aoc-2023/internal/util"
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	util.Init(one, two)
}

func one(f *os.File) {
	seeds, almanac := parseAlmanac(f)

	locations := make([]int, len(seeds))

	for i, seed := range seeds {
		soil := almanac["seed-to-soil"].apply(seed)
		fertilizer := almanac["soil-to-fertilizer"].apply(soil)
		water := almanac["fertilizer-to-water"].apply(fertilizer)
		light := almanac["water-to-light"].apply(water)
		temp := almanac["light-to-temperature"].apply(light)
		humidity := almanac["temperature-to-humidity"].apply(temp)
		loc := almanac["humidity-to-location"].apply(humidity)
		locations[i] = loc
	}
	slices.Sort(locations)
	fmt.Printf("Lowest location: %v\n", locations[0])

}

func parseAlmanac(f *os.File) ([]int, AlmanacMap) {
	var seeds []int

	b := new(bytes.Buffer)
	_, err := b.ReadFrom(f)
	if err != nil {
		panic(err)
	}
	contents := strings.Split(b.String(), "\n")

	seeds = parseDigits(contents[0])
	almanac := parseMaps(contents[1:])

	return seeds, almanac
}

type AlmanacMapEntry struct {
	destRangeStart int
	srcRangeStart  int
	rangeLength    int
}

type AlmanacMap map[string]AlmanacCategoryMap

type AlmanacCategoryMap map[int]int

func (m AlmanacCategoryMap) apply(src int) int {
	dest, ok := m[src]
	if !ok {
		return src
	} else {
		return dest
	}
}

func parseMaps(lines []string) AlmanacMap {

	mapMap := AlmanacMap{
		"seed-to-soil":            make(AlmanacCategoryMap),
		"soil-to-fertilizer":      make(AlmanacCategoryMap),
		"fertilizer-to-water":     make(AlmanacCategoryMap),
		"water-to-light":          make(AlmanacCategoryMap),
		"light-to-temperature":    make(AlmanacCategoryMap),
		"temperature-to-humidity": make(AlmanacCategoryMap),
		"humidity-to-location":    make(AlmanacCategoryMap),
	} // map... map map map, map map.

	// the current map type
	var curr string
	mapHeaderRegex := regexp.MustCompile(`(.+) map:`)
	for _, line := range lines {
		if line == "" {
			// we hit a break between maps, so reset the map type
			curr = ""
			continue
		}
		if curr == "" {
			// we _should_ be on a map header line. Let's make a new entry
			matches := mapHeaderRegex.FindStringSubmatch(line)
			curr = matches[1]
			fmt.Println("Parsing " + curr)
		} else {
			// we've got to parse out the values for the map
			digits := parseDigits(line)
			if len(digits) != 3 {
				panic(errors.New("error: not enough digits for entry"))
			}
			entry := AlmanacMapEntry{
				destRangeStart: digits[0],
				srcRangeStart:  digits[1],
				rangeLength:    digits[2],
			}

			currMap := mapMap[curr]

			for i := entry.rangeLength; i > 0; i-- {
				offset := i - 1
				currMap[entry.srcRangeStart+offset] = entry.destRangeStart + offset
			}
		}
	}
	return mapMap
}

func parseDigits(line string) []int {
	digitsRegex := regexp.MustCompile(`(\d+)`)
	digitStrings := digitsRegex.FindAllString(line, -1)

	digits := make([]int, len(digitStrings))
	for i, v := range digitStrings {
		digit, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		digits[i] = digit
	}
	return digits
}

func two(f *os.File) {

}
