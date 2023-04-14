package valign

import "strings"

func On(search string, lines ...string) (result []string) {
	matches := make(map[int]int) // map[line]at
	max := 0
	for l, line := range lines {
		at := strings.Index(line, search)
		if at >= 0 {
			matches[l] = at
		}
		if at > max {
			max = at
		}
	}
	for l, line := range lines {
		at, ok := matches[l]
		if !ok {
			result = append(result, line)
		} else {
			result = append(result, line[:at]+strings.Repeat(" ", max-at)+line[at:])
		}
	}
	return result
}
