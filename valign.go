package valign

import "strings"

func On(search string, lines ...string) (result []string) {
	matches := make(map[int]int) // map[line]at
	maxLength := 0
	for l, line := range lines {
		at := strings.Index(line, search)
		if at >= 0 {
			matches[l] = at
		}
		if at > maxLength {
			maxLength = at
		}
	}
	for l, line := range lines {
		at, ok := matches[l]
		if !ok {
			result = append(result, line)
		} else {
			result = append(result, line[:at]+strings.Repeat(" ", maxLength-at)+line[at:])
		}
	}
	return result
}

func Blocks(match string, lines ...string) (results [][]string) {
	if len(lines) == 0 {
		return results
	}
	matching := strings.Contains(lines[0], match)
	block := []string{lines[0]}
	for _, line := range lines[1:] {
		if strings.Contains(line, match) != matching {
			results = append(results, block)
			block = make([]string, 0, 0)
			matching = !matching
		}
		block = append(block, line)
	}
	if len(block) > 0 {
		results = append(results, block)
	}
	return results
}
