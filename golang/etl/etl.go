package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := map[string]int{}
	for key, val := range in {
		for _, letter := range val {
			letter = strings.ToLower(letter)
			result[letter] = key
		}
	}
	return result
}
