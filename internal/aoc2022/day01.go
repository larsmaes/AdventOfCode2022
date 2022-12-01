package aoc2022

import "sort"

func Day01part01(input []int) int {

	return elfCalories(input)[0]

}

func Day01part02(input []int) int {

	ec := elfCalories(input)
	return ec[0] + ec[1] + ec[2]

}

func elfCalories(input []int) []int {

	elfCalories := make([]int, 0)

	total := 0
	for _, c := range input {
		total += c
		if c == 0 {
			elfCalories = append(elfCalories, total)
			total = 0
		}
	}
	elfCalories = append(elfCalories, total)
	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))

	return elfCalories
}
