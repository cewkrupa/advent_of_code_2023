package main

import (
	"aoc-2023/internal/util"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	util.Init(one, two)
}

type Game struct {
	Winning []int
	Have    []int
}

func one(file *os.File) {
	games := parseGames(file)
	score := scoreGames(games)
	fmt.Printf("score: %v\n", score)
}

func parseGames(f *os.File) []Game {
	scanner := bufio.NewScanner(f)
	var games []Game
	for scanner.Scan() {
		line := scanner.Text()
		hands := strings.Split(line, ":")
		numbers := strings.Split(hands[1], "|")

		winning := strings.Split(numbers[0], " ")
		have := strings.Split(numbers[1], " ")

		var wInts []int
		for _, wStr := range winning {
			if wStr == "" {
				continue
			}
			w, err := strconv.Atoi(wStr)
			if err != nil {
				panic(err)
			}
			wInts = append(wInts, w)
		}
		var hInts []int
		for _, hStr := range have {
			if hStr == "" {
				continue
			}
			h, err := strconv.Atoi(hStr)
			if err != nil {
				panic(err)
			}
			hInts = append(hInts, h)
		}

		game := Game{wInts, hInts}
		games = append(games, game)

	}

	return games
}

func scoreGames(gs []Game) int {
	var finalScore int

	for i, g := range gs {
		var gameScore int
		var timesMatched int
		for _, num := range g.Have {
			matches := slices.ContainsFunc(g.Winning, func(idx int) bool {
				return idx == num
			})
			if matches {
				if timesMatched < 1 {
					gameScore = 1
				} else {
					gameScore = gameScore * 2
				}
				timesMatched++
			}
		}
		fmt.Printf("game %v: %v\n", i, gameScore)
		finalScore += gameScore
	}
	return finalScore
}

func two(file *os.File) {

}
