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
