package main

import (
	"fmt"
	"os"

	"github.com/larsmaes/AdventOfCode2022/internal/aoc2022"
	"github.com/larsmaes/AdventOfCode2022/internal/utils"
)

func main() {

	input, err := utils.ReadIntFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading input file: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Day1: %d", aoc2022.Day1part1(input))
}
