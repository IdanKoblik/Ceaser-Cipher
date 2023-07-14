package main

func encode(text string, shift rune) string {
	encoded := ""

	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			encoded += string(((char-'A')+shift)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			encoded += string(((char-'a')+shift)%26 + 'a')
		} else {
			encoded += string(char)
		}
	}

	return encoded
}
