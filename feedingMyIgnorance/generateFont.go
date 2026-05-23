package main

import (
	"strings"
)

func GenerateFont() map[rune][]string {
	m := make(map[rune][]string)

	for i := 32; i <= 126; i++ {
		var char []string
		r := rune(i)
		for j := 0; j < 8; j++ {
			var row string
			for k := 0; k < 8; k++ {
				switch {
				case r == ' ':
					row += " "
				default:
					row += "*"
				}
			}
			char = append(char, row)
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
