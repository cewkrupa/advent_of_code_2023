package main

import (
	"aoc-2023/internal/util"
	"bytes"
	"errors"
	"fmt"
	"math"
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

type AlmanacCategoryEntry struct {
	destRangeStart int
	srcRangeStart  int
	rangeLength    int
}

type AlmanacMap map[string]AlmanacCategoryList

func (a AlmanacMap) apply(seed int) int {
	soil := a["seed-to-soil"].apply(seed)
	fertilizer := a["soil-to-fertilizer"].apply(soil)
	water := a["fertilizer-to-water"].apply(fertilizer)
	light := a["water-to-light"].apply(water)
	temp := a["light-to-temperature"].apply(light)
	humidity := a["temperature-to-humidity"].apply(temp)
	loc := a["humidity-to-location"].apply(humidity)
	return loc
}

type AlmanacCategoryList []AlmanacCategoryEntry

func (l AlmanacCategoryList) apply(src int) int {
	var correspondingEntry *AlmanacCategoryEntry
	for _, e := range l {
		e := &e
		if src >= e.srcRangeStart && src < e.srcRangeStart+e.rangeLength {
			correspondingEntry = e
			break
		}
	}

	if correspondingEntry == nil {
		return src
	} else {
		offset := src - correspondingEntry.srcRangeStart
		return correspondingEntry.destRangeStart + offset
	}
}

func parseMaps(lines []string) AlmanacMap {

	mapMap := AlmanacMap{
		"seed-to-soil":            AlmanacCategoryList{},
		"soil-to-fertilizer":      AlmanacCategoryList{},
		"fertilizer-to-water":     AlmanacCategoryList{},
		"water-to-light":          AlmanacCategoryList{},
		"light-to-temperature":    AlmanacCategoryList{},
		"temperature-to-humidity": AlmanacCategoryList{},
		"humidity-to-location":    AlmanacCategoryList{},
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
			entry := AlmanacCategoryEntry{
				destRangeStart: digits[0],
				srcRangeStart:  digits[1],
				rangeLength:    digits[2],
			}

			mapMap[curr] = append(mapMap[curr], entry)
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
	seedRanges, almanac := parseAlmanac(f)

	lowestLoc := math.MaxInt
	for i := 0; i < len(seedRanges); i += 2 {
		start := seedRanges[i]
		length := seedRanges[i+1]

		fmt.Printf("Applying seed range %v, %v\n", start, length)

		chunkNumber := 50

		receiver := make(chan int)
		chunks := getRangeChunks(start, length, chunkNumber)

		for _, c := range chunks {
			go findLowestInChunk(almanac, c.start, c.size, receiver)
		}

		for j := 0; j < len(chunks); j++ {
			select {
			case l := <-receiver:
				if l < lowestLoc {
					lowestLoc = l
				}
			}
		}

	}

	fmt.Printf("Lowest location: %v\n", lowestLoc)
}

type Chunk struct {
	start int
	size  int
}

func getRangeChunks(start int, length int, chunkNumber int) []Chunk {
	var chunks []Chunk

	chunkSize := length / chunkNumber
	var sizeSum int
	for i := 0; i < chunkNumber; i++ {
		chunk := Chunk{
			start: start + sizeSum,
			size:  chunkSize,
		}
		if i == chunkNumber-1 {
			chunk.size = length - sizeSum
		}
		sizeSum += chunk.size
		chunks = append(chunks, chunk)
	}
	if sizeSum != length {
		panic("invalid chunk lengths")
	}
	return chunks
}

func findLowestInChunk(a AlmanacMap, chunkStart int, chunkSize int, out chan<- int) {
	lowest := math.MaxInt
	for i := 0; i < chunkSize; i++ {
		src := chunkStart + i
		loc := a.apply(src)
		if loc < lowest {
			lowest = loc
		}
	}
	out <- lowest
}
