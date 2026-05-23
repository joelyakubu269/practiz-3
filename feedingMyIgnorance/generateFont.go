package main

func GenerateFont() map[rune][]string {
	m := make(map[rune][]string)
	var char []string
	for i := 32; i <= 126; i++ {
		r := rune(i)
		for j := 0; j < 8; j++ {
			var row string
			for k := 0; k < 8; k++ {
				switch {
				case r == ' ':

				}
			}
		}
	}

}
