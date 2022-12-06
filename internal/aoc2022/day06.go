package aoc2022

func Day06Part01(input string) int {
	window := 4
	var marker int

	for i := 0; i <= len(input)-window; i++ {
		packet := input[i : window+i]

		if isMarker(packet) {
			marker = i + window
			break
		}

	}

	return marker
}

func Day06Part02(input string) int {
	window := 14
	var marker int

	for i := 0; i <= len(input)-window; i++ {
		packet := input[i : window+i]

		if isMarker(packet) {
			marker = i + window
			break
		}

	}

	return marker
}

func isMarker(packet string) bool {
	found := false

	for i, c := range packet {
		for i2, c2 := range packet {
			if i != i2 {
				if c == c2 {
					found = true
				}
			}
		}
	}

	return !found
}
