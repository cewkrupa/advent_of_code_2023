package main

import (
	"aoc-2023/internal/util"
	"fmt"
)

func main() {
	util.Init(one, two)
}

func one(fileName string) {
	fmt.Println("one")
}

func two(fileName string) {
	fmt.Println("two")
}
