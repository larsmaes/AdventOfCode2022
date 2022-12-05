package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

	var stateLines []string
	var directiveLines []string
	reader := "state"
	for _, line := range input {
		switch reader {
		case "state":
			if line != "" {
				stateLines = append(stateLines, line)
			} else {
				reader = "directive"
			}
		case "directive":
			directiveLines = append(directiveLines, line)
		}
	}

	state := parseStatelines(stateLines)
	directives := parseDirectiveLines(directiveLines)

	fmt.Printf("Day5 part 1: %s\n", aoc2022.Day05Part01(state, directives))

	state = parseStatelines(stateLines)
	directives = parseDirectiveLines(directiveLines)

	fmt.Printf("Day5 part 2: %s\n", aoc2022.Day05Part02(state, directives))
}

func parseStatelines(lines []string) map[int][]string {
	stateMap := make(map[int][]string)
	for i := len(lines) - 1; i >= 0; i-- {
		if i != len(lines)-1 {
			for o := 0; o <= 8; o++ {
				offset := 4*o + 1
				if string(lines[i][offset]) != " " {
					stateMap[o] = append(stateMap[o], string(lines[i][offset]))
				}
			}
		}
	}
	return stateMap
}

func parseDirectiveLines(lines []string) []aoc2022.Directive {
	var directives []aoc2022.Directive
	for _, l := range lines {
		words := strings.Split(l, " ")
		amount, _ := strconv.Atoi(words[1])
		from, _ := strconv.Atoi(words[3])
		to, _ := strconv.Atoi(words[5])
		directives = append(directives, aoc2022.Directive{
			Amount: amount,
			From:   from - 1,
			To:     to - 1,
		})
	}
	return directives
}
