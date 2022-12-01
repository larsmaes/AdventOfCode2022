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

	fmt.Printf("Day1 part 1: %d", aoc2022.Day01part01(input))
	fmt.Printf("Day1 part 2: %d", aoc2022.Day01part02(input))
}
