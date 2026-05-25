package main

import "strings"

const size = 8

func GenerateFont() map[rune][]string {
	font := map[rune][]string{}
	for i := 32; i <= 126; i++ {
		for rows := 0; rows < size; rows++ {
			var lines strings.Builder
			for cols := 0; cols < size; cols++ {

			}
		}
	}
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
	return strings.ContainsRune("AEIOUaeiou", r)
}

func pixel(r rune, rows, cols int) rune {
	switch classify(r) {
	case "space":
		return ' '
	case "upper":
		if rows == 0 || rows == 7 || cols == 0 || cols == 7 {
			return '*'
		}
		if cols == 2 || cols == 5 {
			return '*'
		}
		if rows == 3 {
			return '*'
		}
		if (rows == cols) || (rows+cols == 7) {
			return '*'
		}
		return '.'
	case "lower":
		top := 2
		bottom := 6
		if isAscender(r) {
			top = 1
		}
		if isDescender(r) {
			bottom = 7
		}
		if cols == 3 {
			return '*'
		}
		if (rows == top || rows == bottom) && cols >= 2 && cols <= 5 {
			return '*'
		}
		if isVowel(r) && (rows == top || rows == bottom) {
			return '*'
		}
		return '.'
	case "digit":
		if rows == 0 || rows == 7 {
			return '*'
		}
		if cols == 3 { // to give it a vertical spine
			return '*'
		}
		if (rows+cols)%4 == 0 { // for uniqueness
			return '*'
		}
		return '.'
	case "symbol":
		if rows == 0 || rows == 7 || cols == 0 || cols == 7 {
			return '*'
		}
		if (rows+cols+int(r))%3 == 0 {
			return '*'
		}
		return '*'
	}
	return '.'
}
