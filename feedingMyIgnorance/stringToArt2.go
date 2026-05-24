package main

import "strings"

func StringToArt(input string) string {
	for _, r := range input {
		if (r < '0' || r > '9') && r != '\n' {
			return ""
		}
	}
	m := map[rune][]string{
		'0': {
			" ___ ",
			" | | ",
			" | | ",
			" | | ",
			" ___ ",
		},
		'1': {
			"  /  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"     ",
		},

		'2': {
			" ___ ",
			"    |",
			" ___ ",
			"|    ",
			"_____",
		},

		'3': {
			"____ ",
			"    |",
			" ___ ",
			"    |",
			"____ ",
		},

		'4': {
			"|   |",
			"|___|",
			"    |",
			"    |",
			"    |",
		},

		'5': {
			"_____",
			"|    ",
			"_____",
			"    |",
			"_____",
		},

		'6': {
			" ____",
			"|    ",
			"_____",
			"|   |",
			" ___|",
		},

		'7': {
			"_____",
			"    |",
			"   / ",
			"  /  ",
			" /   ",
		},

		'8': {
			" ___ ",
			"|   |",
			" ___ ",
			"|   |",
			" ___ ",
		},

		'9': {
			" ___ ",
			"|   |",
			" ____",
			"    |",
			" ___|",
		},
	}
	parts := strings.Split(input, "\n")
	var result strings.Builder
	for _, r := range parts {
		for i := 0; i < 5; i++ {

			for _, c := range r {
				val, ok := m[c]
				if !ok {
					return ""
				}
				result.WriteString(val[i])

			}
			result.WriteString("\n")
		}

	}
	return result.String()
}
