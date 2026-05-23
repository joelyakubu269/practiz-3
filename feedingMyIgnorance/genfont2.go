package main

import (
	"strings"
)

func GenerateFont() map[rune][]string {
	m := make(map[rune][]string)
	for i := 32; i <= 126; i++ {
		char := make([]string, 8)
		r := rune(i)
		for j := 0; j < 8; j++ {
			var row strings.Builder
			for k := 0; k < 8; k++ {
				row.WriteRune(pixel(r, j, k))
			}
			char[j] = row.String()
		}
		m[r] = char
	}
	return m
}

func classify(r rune) string {
	switch {
	case r >= 'A' && r <= 'Z':
		return "upper"
	case r >= 'a' && r <= 'z':
		return "lower"
	case r == ' ':
		return "space"
	case r >= '0' && r <= '9':
		return "digit"
	default:
		return "symbol"
	}
}

func isAscender(r rune) bool {
	return strings.ContainsRune("bdfhklt", r)
}

func isDescender(r rune) bool {
	return strings.ContainsRune("gjpqy", r)
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouAEIOU", r)
}

func pixel(r rune, row, col int) rune {
	// Seed variation per character so each one looks distinct
	n := int(r)

	switch classify(r) {
	case "space":
		return ' '

	case "upper":
		// Outline box: always draw top, bottom, left, right edges
		topRow, botRow := 1, 6
		leftCol, rightCol := 1, 6
		if row == topRow || row == botRow {
			if col >= leftCol && col <= rightCol {
				return '*'
			}
		}
		if col == leftCol || col == rightCol {
			if row >= topRow && row <= botRow {
				return '*'
			}
		}
		// Midbar: vowels and every 3rd letter get a horizontal bar
		if (isVowel(r) || n%3 == 0) && row == 3 && col >= leftCol && col <= rightCol {
			return '*'
		}
		// Diagonal accent: seeded by char value gives each letter a unique inner mark
		if row-topRow == (col-leftCol+n)%(botRow-topRow+1) {
			return '*'
		}

	case "lower":
		asc, desc := isAscender(r), isDescender(r)
		// x-height zone rows 3-5; ascender adds rows 1-2; descender adds rows 6-7
		top := 3
		if asc {
			top = 1
		}
		bot := 5
		if desc {
			bot = 7
		}
		// Outline
		if (row == top || row == bot) && col >= 1 && col <= 6 {
			return '*'
		}
		if (col == 1 || col == 6) && row >= top && row <= bot {
			return '*'
		}
		// Per-character inner detail
		mid := (top + bot) / 2
		if row == mid && col >= 2 && col <= 5 && n%2 == 0 {
			return '*'
		}

	case "digit":
		// Segments: top/bottom bars + varying vertical sides
		topRow, botRow := 1, 6
		if (row == topRow || row == botRow) && col >= 1 && col <= 6 {
			return '*'
		}
		// Left side on upper half if bit 0 set, right side if bit 1 set, etc.
		mid := 3
		if row == mid && col >= 1 && col <= 6 && (n-'0')%2 == 0 {
			return '*'
		}
		upperLeft := (n-'0')&1 != 0
		upperRight := (n-'0')&2 != 0
		lowerLeft := (n-'0')&4 != 0
		lowerRight := (n-'0')&8 != 0 // only 0-9 so max 9 = 1001
		if row > topRow && row < mid && col == 1 && upperLeft {
			return '*'
		}
		if row > topRow && row < mid && col == 6 && upperRight {
			return '*'
		}
		if row > mid && row < botRow && col == 1 && lowerLeft {
			return '*'
		}
		if row > mid && row < botRow && col == 6 && lowerRight {
			return '*'
		}

	case "symbol":
		// Checkerboard seeded by char value — each symbol gets a unique phase
		if (row+col+n)%2 == 0 && row > 0 && row < 7 && col > 0 && col < 7 {
			return '*'
		}
		// Always draw a border so symbols have a frame
		if (row == 1 || row == 6) && col >= 1 && col <= 6 {
			return '*'
		}
		if (col == 1 || col == 6) && row >= 1 && row <= 6 {
			return '*'
		}
	}

	return '.'
}

// func main() {
// 	font := GenerateFont()
// 	// Print a sample side by side
// 	sample := "Hello 123"
// 	bufs := make([]strings.Builder, 8)
// 	for _, r := range sample {
// 		g := font[r]
// 		for i, line := range g {
// 			bufs[i].WriteString(line)
// 			bufs[i].WriteByte(' ')
// 		}
// 	}
// 	fmt.Println("Sample:", sample)
// 	for _, b := range bufs {
// 		fmt.Println(b.String())
// 	}
// }
