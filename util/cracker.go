package util

func Crack(text string) string {
	var decoded string

	for shift := 1; shift <= 25; shift++ {
		for _, char := range text {
			if char >= 'A' && char <= 'Z' {
				decoded += string((char-'A'-rune(shift)+26)%26 + 'A')
			} else if char >= 'a' && char <= 'z' {
				decoded += string((char-'a'-rune(shift)+26)%26 + 'a')
			} else {
				decoded += string(char)
			}
		}
		decoded += "\n"
	}

	return decoded
}
