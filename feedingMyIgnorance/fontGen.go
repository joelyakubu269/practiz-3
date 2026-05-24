package main

func GenerateFont() map[rune][]string {

}
func classify(c rune) string {
	if r >= '0' && r <= '9' {
		return "digit"
	}
	if r >= 'A' && r < 'Z' {
		return "upper"
	}
	if r >= 'a' && r < 'z' {
		if isAscender(r) {
			return "ascender"
		}else if isDescender(r) {
			return descender
		}else {
			return "lower"
		}
	}else {
		retrun "symbol"
	}
}
func isAscender(c rune) bool{
	return strings.ContainsRune("bdfhklt",c)
}
func isDescender(c rune) bool {
	return strings.ContainsRune("gjpqy",c)
}
func isVowel(c rune) bool {
	return strings.ContainsRune("AEIOUaeiou",c)
}
