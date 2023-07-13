package main

import (
	"fmt"
)

func encoder(text string, shift rune) string {
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

func decoder(text string, shift rune) string {
	decoded := ""

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

func main() {
	for true {
		fmt.Println("Welcome to Caesar Cipher decoder/encoder in Golang")

		fmt.Println("Please choose an option:")
		fmt.Println("1. Encode")
		fmt.Println("2. Decode")
		fmt.Println("3. Exit")

		var option int
		_, _ = fmt.Scanln(&option)

		if option == 1 {
			fmt.Print("Please enter the text to encode: ")
			var text string
			_, _ = fmt.Scan(&text)

			fmt.Print("Please enter the shift: ")
			var shift rune
			_, _ = fmt.Scan(&shift)

			fmt.Println(encoder(text, shift))
		} else if option == 2 {
			fmt.Print("Please enter the text to decode: ")
			var text string
			_, _ = fmt.Scan(&text)

			fmt.Print("Please enter the shift: ")
			var shift rune
			_, _ = fmt.Scanln(&shift)

			fmt.Println(decoder(text, shift))
		} else if option == 3 {
			return
		} else {
			fmt.Println("Invalid option")
			return
		}
	}
}
