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
	ID        int
	Winning   []int
	Have      []int
	Score     int
	Matches   int
	Instances int
}

func (g Game) String() string {
	return fmt.Sprintf("Game %v - Instances: %v Matches: %v Score: %v\n", g.ID, g.Instances, g.Matches, g.Score)
}

func one(file *os.File) {
	games := parseGames(file)
	score := scoreGames(games)
	fmt.Printf("score: %v\n", score)
}

func parseGames(f *os.File) []Game {
	scanner := bufio.NewScanner(f)
	var games []Game
	gameNumber := 1
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

		game := Game{ID: gameNumber, Winning: wInts, Have: hInts, Instances: 1}
		games = append(games, game)
		gameNumber++

	}

	return games
}

func scoreGames(gList []Game) int {
	var finalScore int

	for i, g := range gList {
		scoreGame(&g)
		finalScore += g.Score
		gList[i] = g
	}
	return finalScore
}

func scoreGame(g *Game) {
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
	g.Matches = timesMatched
	g.Score = gameScore
}

func two(file *os.File) {
	games := parseGames(file)
	scoreGames(games)
	totalCards := processGames(games)
	fmt.Printf("totalCards: %v\n", totalCards)
}

func processGames(gameList []Game) int {
	tableLength := len(gameList)
	var totalCards int

	for i, g := range gameList {
		totalCards += g.Instances
		if g.Matches > 0 {
			for j := 1; j <= g.Matches && j+i < tableLength; j++ {
				nextIdx := i + j
				next := &gameList[nextIdx]

				next.Instances = next.Instances + g.Instances
			}
		}
	}
	return totalCards
}
