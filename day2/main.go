package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
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

func one(input string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("part one")

	// define bag start
	var possibleIdCount int
	boardConfig := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		r1 := regexp.MustCompile(`Game (\d+):(.+)`)
		sections := r1.FindAllStringSubmatch(line, -1)

		gameID, err := strconv.Atoi(sections[0][1])
		if err != nil {
			panic(err)
		}

		/*
			"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
			-> ["3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"]
		*/
		turnStrings := strings.Split(sections[0][2], ";")
		possible := true
		for _, t := range turnStrings {
			t := strings.Trim(t, " ")
			/*
				"3 blue, 4 red"
				-> ["3 blue", "4 red"]
			*/
			cubeSets := strings.Split(t, ",")

			setCounts := map[string]int{
				"blue":  0,
				"red":   0,
				"green": 0,
			}

			for _, cs := range cubeSets {
				/*
					"3 blue"
					-> ["3", "blue"]
				*/
				cs := strings.Trim(cs, " ")
				set := strings.Split(cs, " ")

				if len(set) != 2 {
					// we did something wrong in parsing...
					err := fmt.Errorf("error: Cube set not 2 parts: %v", set)
					panic(err)
				}
				cubeNumber, err := strconv.Atoi(set[0])
				if err != nil {
					panic(err)
				}
				cubeColor := set[1]

				setCounts[cubeColor] = cubeNumber
			}

			for color, maxCount := range boardConfig {
				if setCounts[color] > maxCount {
					possible = false
				}
			}
		}
		// add to sum
		fmt.Printf("Game ID %v possible: %v\n", gameID, possible)
		if possible {
			possibleIdCount += gameID
		}
	}
	fmt.Printf("ID Sum: %v", possibleIdCount)
}

type Game struct {
	ID    int
	Turns []TurnSet
}
type TurnSet struct {
	Red   int
	Blue  int
	Green int
}

func two(input string) {
	fmt.Println("part two")
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	var games []Game

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		games = append(games, *game)

	}

	var powerSum int
	for _, g := range games {
		minSet := TurnSet{}
		for _, ts := range g.Turns {
			minSet.Red = max(minSet.Red, ts.Red)
			minSet.Blue = max(minSet.Blue, ts.Blue)
			minSet.Green = max(minSet.Green, ts.Green)
		}
		minSetPower := minSet.Red * minSet.Blue * minSet.Green
		fmt.Printf("power: %v\n", minSetPower)
		powerSum += minSetPower
	}
	fmt.Printf("powerSum: %v\n", powerSum)

}
func parseGame(gameString string) *Game {
	game := Game{}
	r1 := regexp.MustCompile(`Game (\d+):(.+)`)
	sections := r1.FindAllStringSubmatch(gameString, -1)

	gameID, err := strconv.Atoi(sections[0][1])
	if err != nil {
		panic(err)
	}
	game.ID = gameID

	turnStrings := strings.Split(sections[0][2], ";")
	for _, t := range turnStrings {
		ts := TurnSet{}

		t := strings.Trim(t, " ")

		cubeSets := strings.Split(t, ",")

		setCounts := map[string]int{
			"blue":  0,
			"red":   0,
			"green": 0,
		}

		for _, cs := range cubeSets {
			cs := strings.Trim(cs, " ")
			set := strings.Split(cs, " ")

			if len(set) != 2 {
				// we did something wrong in parsing...
				err := fmt.Errorf("error: Cube set not 2 parts: %v", set)
				panic(err)
			}
			cubeNumber, err := strconv.Atoi(set[0])
			if err != nil {
				panic(err)
			}
			cubeColor := set[1]

			setCounts[cubeColor] = cubeNumber
		}
		ts.Red = setCounts["red"]
		ts.Blue = setCounts["blue"]
		ts.Green = setCounts["green"]
		game.Turns = append(game.Turns, ts)
	}
	return &game
}
