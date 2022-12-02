package aoc2022

import "strings"

var shapeScores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var gameScores = map[string]int{
	"win":  6,
	"lose": 0,
	"draw": 3,
}

func Day02part01(input []string) int {
	totalScore := 0

	for _, game := range input {
		shapes := strings.Split(game, " ")
		totalScore += determineShapeScore(shapes[1])
		totalScore += determineGameScore(shapes[0], shapes[1])
	}
	return totalScore

}

func Day02part02(input []string) int {

	totalScore := 0

	for _, game := range input {
		shapes := strings.Split(game, " ")
		totalScore += determineGameOutcomeScore(shapes[0], shapes[1])
	}
	return totalScore

}

func determineShapeScore(shape string) int {
	return shapeScores[shape]
}

func determineGameScore(hisShape, myShape string) int {
	switch hisShape {
	case "A":
		switch myShape {
		case "X":
			return gameScores["draw"]
		case "Y":
			return gameScores["win"]
		case "Z":
			return gameScores["lose"]
		}
	case "B":
		switch myShape {
		case "X":
			return gameScores["lose"]
		case "Y":
			return gameScores["draw"]
		case "Z":
			return gameScores["win"]
		}
	case "C":
		switch myShape {
		case "X":
			return gameScores["win"]
		case "Y":
			return gameScores["lose"]
		case "Z":
			return gameScores["draw"]
		}

	}
	return 0
}

func determineGameOutcomeScore(hisShape, desiredOutcome string) int {
	switch hisShape {
	case "A": // rock
		switch desiredOutcome {
		case "X": //lose
			return shapeScores["Z"] + gameScores["lose"]
		case "Y": // draw
			return shapeScores["X"] + gameScores["draw"]
		case "Z": //win
			return shapeScores["Y"] + gameScores["win"]
		}
	case "B": // paper
		switch desiredOutcome {
		case "X": //lose
			return shapeScores["X"] + gameScores["lose"]
		case "Y": // draw
			return shapeScores["Y"] + gameScores["draw"]
		case "Z": //win
			return shapeScores["Z"] + gameScores["win"]
		}
	case "C": // scissors
		switch desiredOutcome {
		case "X": //lose
			return shapeScores["Y"] + gameScores["lose"]
		case "Y": // draw
			return shapeScores["Z"] + gameScores["draw"]
		case "Z": //win
			return shapeScores["X"] + gameScores["win"]
		}
	}
	return 0
}
