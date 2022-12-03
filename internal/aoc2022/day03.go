package aoc2022

import (
	"github.com/juliangruber/go-intersect"
)

func Day03part01(input []string) int {

	var sum int
	for _, rucksack := range input {

		c1 := []rune(rucksack[:len(rucksack)/2])
		c2 := []rune(rucksack[len(rucksack)/2:])

		i := intersect.Hash(c1, c2)
		item := i[0].(rune)
		if item > 95 {
			sum += int(item - 96)
		} else {
			sum += int(item - 38)
		}

	}

	return sum
}

func Day03part02(input []string) int {
	var rucksackGroups [][]string
	var group []string
	i := 0

	for _, rucksack := range input {
		group = append(group, rucksack)
		if i > 1 {
			rucksackGroups = append(rucksackGroups, group)
			group = []string{}
			i = 0
		} else {
			i++
		}

	}

	sum := 0
	for _, g := range rucksackGroups {
		r1 := g[0]
		r2 := g[1]
		r3 := g[2]

		i1 := intersect.Hash(r1, r2)
		i2 := intersect.Hash(i1, r3)

		item := i2[0].(uint8)
		if item > 95 {
			sum += int(item - 96)
		} else {
			sum += int(item - 38)
		}

	}
	return sum
}
