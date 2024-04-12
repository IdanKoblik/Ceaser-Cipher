package util

func Decode(text string, shift rune) string {
	var decoded = ""

	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			decoded += string((char-'A'-shift+26)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			decoded += string((char-'a'-shift+26)%26 + 'a')
		} else {
			decoded += string(char)
		}
	}

	return decoded
}
