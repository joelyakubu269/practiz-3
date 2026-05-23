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
			var row string
			for k := 0; k < 8; k++ {
				row += string(pixel(r, j, k))
			}
			char[j] = row
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
func rowZone(row int) string {
	switch {
	case row <= 1:
		return "top"
	case row <= 3:
		return "upper"
	case row == 4:
		return "middle"
	case row <= 6:
		return "lower"
	default:
		return "bottom"
	}
}
func pixel(r rune, row, col int) rune {
	class := classify(r)
	zone := rowZone(row)

	switch class {
	case "space":
		return ' '
	case "digit":
		if row == col || col == 0 || col == 7 {
			return '*'
		}
	case "upper":
		if zone == "top" || zone == "upper" {
			return '*'
		}
	case "lower":
		if zone == "lower" || zone == "bottom" {
			return '*'
		}
	case "symbol":
		if (row+col)%2 == 0 {
			return '*'
		}
	}
	return '.'
}

// func main() {
// 	font := GenerateFont()

// 	for r, char := range font {

// 		fmt.Printf("%c\n", r)

// 		for _, line := range char {
// 			fmt.Println(line)
// 		}

// 		fmt.Println()
// 	}
// }
