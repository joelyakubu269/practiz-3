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
func isAscender(r rune) bool {
	return strings.ContainsRune("bdfhklt", r)
}
func isDescender(r rune) bool {
	return strings.ContainsRune("gjpqy", r)
}
func isVowel(r rune) bool {
	return strings.ContainsRune("AEIOUaeiou", r)
}
