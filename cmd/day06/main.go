package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/larsmaes/AdventOfCode2022/internal/aoc2022"
)

func main() {

	bytesRead, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	file_content := string(bytesRead)
	input := strings.Split(file_content, "\n")

	fmt.Printf("Day6 part 1: %d\n", aoc2022.Day06Part01(input[0]))
	fmt.Printf("Day6 part 2: %d\n", aoc2022.Day06Part02(input[0]))
}
