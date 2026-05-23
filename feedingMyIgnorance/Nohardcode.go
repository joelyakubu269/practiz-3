package main

import (
	"fmt"
	"strings"
)

const size = 8

// ---------------- FONT GENERATOR ----------------

func GenerateFont() map[rune][]string {
	font := make(map[rune][]string)

	for i := 32; i <= 126; i++ {
		r := rune(i)
		glyph := make([]string, size)

		for row := 0; row < size; row++ {
			var b strings.Builder

			for col := 0; col < size; col++ {
				b.WriteRune(pixel(r, row, col))
			}

			glyph[row] = b.String()
		}

		font[r] = glyph
	}

	return font
}

// ---------------- CLASSIFICATION ----------------

func classify(r rune) string {
	switch {
	case r >= 'A' && r <= 'Z':
		return "upper"
	case r >= 'a' && r <= 'z':
		return "lower"
	case r >= '0' && r <= '9':
		return "digit"
	case r == ' ':
		return "space"
	default:
		return "symbol"
	}
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouAEIOU", r)
}

func isAscender(r rune) bool {
	return strings.ContainsRune("bdfhklt", r)
}

func isDescender(r rune) bool {
	return strings.ContainsRune("gjpqy", r)
}

// ---------------- PIXEL ENGINE ----------------

func pixel(r rune, row, col int) rune {

	switch classify(r) {

	case "space":
		return ' '

	// ---------------- UPPERCASE ----------------
	case "upper":

		if row == 0 || row == 7 || col == 0 || col == 7 {
			return '*'
		}

		if col == 2 || col == 5 {
			return '*'
		}

		if row == 3 {
			return '*'
		}

		if row == col || row+col == 7 {
			return '*'
		}

		return '.'

	// ---------------- LOWERCASE ----------------
	case "lower":

		top := 2
		bottom := 6

		if isAscender(r) {
			top = 1
		}
		if isDescender(r) {
			bottom = 7
		}

		if col == 3 {
			return '*'
		}

		if (row == top || row == bottom) && col >= 2 && col <= 5 {
			return '*'
		}

		if isVowel(r) && (row == top || row == bottom) {
			return '*'
		}

		return '.'

	// ---------------- DIGITS ----------------
	case "digit":

		if row == 0 || row == 7 {
			return '*'
		}

		if col == 3 {
			return '*'
		}

		if (row+col)%4 == 0 {
			return '*'
		}

		return '.'

	// ---------------- SYMBOLS ----------------
	case "symbol":

		if row == 0 || row == 7 || col == 0 || col == 7 {
			return '*'
		}

		if (row+col+int(r))%3 == 0 {
			return '*'
		}

		return '.'
	}

	return '.'
}

// ---------------- TEST MAIN ----------------

func main() {
	font := GenerateFont()

	// Test a sample string
	sample := "Hello 123!"

	lines := make([]strings.Builder, size)

	for _, r := range sample {
		glyph := font[r]

		for i := 0; i < size; i++ {
			lines[i].WriteString(glyph[i])
			lines[i].WriteString(" ") // spacing between chars
		}
	}

	fmt.Println("INPUT:", sample)
	fmt.Println()

	for i := 0; i < size; i++ {
		fmt.Println(lines[i].String())
	}
}
