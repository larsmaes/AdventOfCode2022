package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntFile(filename string) ([]int, error) {

	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(r)

	var input []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		input = append(input, i)
	}

	return input, nil
}
