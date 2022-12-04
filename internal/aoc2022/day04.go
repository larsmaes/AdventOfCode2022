package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

func Day04Part01(input []string) int {
	fullyContaining := 0
	for _, pair := range input {
		elves := strings.Split(pair, ",")

		elf1range := rangeString2IntString(elves[0])
		elf2range := rangeString2IntString(elves[1])

		if strings.Contains(elf1range, elf2range) {
			fullyContaining++

		} else if strings.Contains(elf2range, elf1range) {
			fullyContaining++
		}
	}

	return fullyContaining
}

func Day04Part02(input []string) int {
	partiallyContaining := 0
	for _, pair := range input {
		elves := strings.Split(pair, ",")

		elf1range := rangeString2IntSlice(elves[0])
		elf2range := rangeString2IntSlice(elves[1])

	found:
		for _, v := range elf1range {
			for _, v2 := range elf2range {
				if v == v2 {
					partiallyContaining++
					break found
				}
			}
		}
	}
	return partiallyContaining
}

func rangeString2IntString(r string) string {
	ranges := strings.Split(r, "-")
	outSlice := []string{}
	low, _ := strconv.Atoi(ranges[0])
	high, _ := strconv.Atoi(ranges[1])
	for i := low; i <= high; i++ {
		outSlice = append(outSlice, fmt.Sprintf("%02d", i))
	}
	return strings.Join(outSlice, ",")
}

func rangeString2IntSlice(r string) []int {
	ranges := strings.Split(r, "-")
	outSlice := []int{}
	low, _ := strconv.Atoi(ranges[0])
	high, _ := strconv.Atoi(ranges[1])
	for i := low; i <= high; i++ {
		outSlice = append(outSlice, i)
	}
	return outSlice
}
