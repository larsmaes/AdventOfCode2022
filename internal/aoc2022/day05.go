package aoc2022

type Directive struct {
	Amount int
	From   int
	To     int
}

func Day05Part01(state map[int][]string, directives []Directive) string {

	for _, directive := range directives {
		for i := 0; i < directive.Amount; i++ {
			var elem string
			elem, state[directive.From] = state[directive.From][len(state[directive.From])-1], state[directive.From][:len(state[directive.From])-1]

			state[directive.To] = append(state[directive.To], elem)

		}
	}

	var outputString string

	for i := 0; i <= len(state); i++ {
		if len(state[i]) > 0 {
			outputString += state[i][len(state[i])-1]
		}
	}
	return outputString
}

func Day05Part02(state map[int][]string, directives []Directive) string {

	for _, directive := range directives {
		var elem []string
		for i := 0; i < directive.Amount; i++ {

			elem, state[directive.From] = append(elem, state[directive.From][len(state[directive.From])-1]), state[directive.From][:len(state[directive.From])-1]

		}
		reverse(elem)
		state[directive.To] = append(state[directive.To], elem...)
	}

	var outputString string

	for i := 0; i <= len(state); i++ {
		if len(state[i]) > 0 {
			outputString += state[i][len(state[i])-1]
		}
	}
	return outputString
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
