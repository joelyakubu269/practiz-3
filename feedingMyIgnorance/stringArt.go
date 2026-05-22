package main

import (
	"strings"
)

func StringToArt(input string) string {
	if len(input) == 0 {
		return ""
	}
	for _, r := range input {
		if (r < '0' || r > '9') && r != '\n' {
			return ""
		}
	}
	m := map[rune][]string{
		'0': {
			" --- ",
			" | | ",
			" | | ",
			" | | ",
			" --- ",
		},
		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			" --- ",
			"    |",
			" --- ",
			"|    ",
			" --- ",
		},

		'3': {
			" --- ",
			"    |",
			" --- ",
			"    |",
			" --- ",
		},

		'4': {
			"|  | ",
			"|  | ",
			" --- ",
			"    |",
			"    |",
		},

		'5': {
			" --- ",
			"|    ",
			" --- ",
			"    |",
			" --- ",
		},

		'6': {
			" --- ",
			"|    ",
			" --- ",
			"|   |",
			" --- ",
		},

		'7': {
			" --- ",
			"    |",
			"   | ",
			"  |  ",
			" |   ",
		},

		'8': {
			" --- ",
			"|   |",
			" --- ",
			"|   |",
			" --- ",
		},

		'9': {
			" --- ",
			"|   |",
			" --- ",
			"    |",
			" --- ",
		},
	}
	words := strings.Split(input, "\n")
	var output string

	for _, r := range words {
		for i := 0; i < 5; i++ {
			for _, c := range r {
				if val, ok := m[c]; ok {
					output += val[i]
				}
			}
			output += "\n"

		}
	}
	return output
}

// func main() {
// 	fmt.Print(StringToArt("12\n2"))
// }
