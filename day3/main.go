package main

import (
	"aoc-2023/internal/util"
	"bufio"
	"fmt"
	"os"
)

func main() {
	util.Init(one, two)
}

type Coordinate struct {
	X int
	Y int
}

type NumberEntry struct {
	val    int
	isPart bool
}

func one(fileName string) {
	fmt.Println("one")

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	numMap, symbolMap := buildMap(file)

	var parts []*NumberEntry
	// now, for every symbol we found, see what numbers are around it and mark them so
	for c := range symbolMap {
		north := numMap[Coordinate{
			X: c.X,
			Y: c.Y - 1,
		}]
		northeast := numMap[Coordinate{
			X: c.X + 1,
			Y: c.Y - 1,
		}]
		east := numMap[Coordinate{
			X: c.X + 1,
			Y: c.Y,
		}]
		southeast := numMap[Coordinate{
			X: c.X + 1,
			Y: c.Y + 1,
		}]
		south := numMap[Coordinate{
			X: c.X,
			Y: c.Y + 1,
		}]
		southwest := numMap[Coordinate{
			X: c.X - 1,
			Y: c.Y + 1,
		}]
		west := numMap[Coordinate{
			X: c.X - 1,
			Y: c.Y,
		}]
		northwest := numMap[Coordinate{
			X: c.X - 1,
			Y: c.Y - 1,
		}]

		neighbors := []*NumberEntry{
			north, northeast, east, southeast, south, southwest, west, northwest,
		}

		for _, n := range neighbors {
			if n != nil && !n.isPart {
				n.isPart = true
				parts = append(parts, n)
			}
		}
	}

	var sum int
	for _, part := range parts {
		sum += part.val
	}

	fmt.Printf("sum: %v\n", sum)
}

func buildMap(file *os.File) (map[Coordinate]*NumberEntry, map[Coordinate]string) {
	numMap := map[Coordinate]*NumberEntry{}
	symbolMap := map[Coordinate]string{}
	scanner := bufio.NewScanner(file)
	var lineNo int
	for scanner.Scan() {
		line := scanner.Text()

		// walk the line
		for i, c := range line {
			coord := Coordinate{
				X: i,
				Y: lineNo,
			}
			cInt := c - '0'
			if cInt <= 9 && cInt >= 0 {
				// we've got a number!
				// see if there's a number in the preceding coordinate
				// if there is, concatenate this one to it,
				// and make a new entry for this one, pointing to the same entry.
				// otherwise, just make a new entry for this one.
				prev := numMap[Coordinate{X: coord.X - 1, Y: lineNo}]
				if prev != nil {
					prev.val = (prev.val * 10) + int(cInt)
					numMap[coord] = prev
				} else {
					numMap[coord] = &NumberEntry{
						val: int(cInt),
					}
				}
			} else if c == '.' {
				// we have a period, continue
			} else {
				// we have some sort of symbol
				symbolMap[coord] = string(c)
			}
		}
		lineNo++
	}
	return numMap, symbolMap
}

func two(fileName string) {
	fmt.Println("two")
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	numMap, symbolMap := buildMap(file)

	var gearSum int
	// now, for every symbol we found, see what numbers are around it and mark them so
	for c, s := range symbolMap {
		north := numMap[Coordinate{
			X: c.X,
			Y: c.Y - 1,
		}]
		northeast := numMap[Coordinate{
			X: c.X + 1,
			Y: c.Y - 1,
		}]
		east := numMap[Coordinate{
			X: c.X + 1,
			Y: c.Y,
		}]
		southeast := numMap[Coordinate{
			X: c.X + 1,
			Y: c.Y + 1,
		}]
		south := numMap[Coordinate{
			X: c.X,
			Y: c.Y + 1,
		}]
		southwest := numMap[Coordinate{
			X: c.X - 1,
			Y: c.Y + 1,
		}]
		west := numMap[Coordinate{
			X: c.X - 1,
			Y: c.Y,
		}]
		northwest := numMap[Coordinate{
			X: c.X - 1,
			Y: c.Y - 1,
		}]

		neighbors := []*NumberEntry{
			north, northeast, east, southeast, south, southwest, west, northwest,
		}

		var neighborParts []*NumberEntry
		for _, n := range neighbors {
			if n != nil && !n.isPart {
				n.isPart = true
				neighborParts = append(neighborParts, n)
			}
		}
		if s == "*" && len(neighborParts) == 2 {
			gearSum += neighborParts[0].val * neighborParts[1].val
		}
	}

	fmt.Printf("gear sum: %v\n", gearSum)
}
