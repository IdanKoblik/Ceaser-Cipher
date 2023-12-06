package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("Welcome to Caesar Cipher decoder/encoder in Golang")

		fmt.Println("Please choose an option:")
		fmt.Println("1. Encode")
		fmt.Println("2. Decode")
		fmt.Println("3. Exit")

		var option int
		_, _ = fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Print("Please enter the text to encode: ")
			var text string
			_, _ = fmt.Scan(&text)

			fmt.Print("Please enter the shift: ")
			var shift rune
			_, _ = fmt.Scan(&shift)

			fmt.Println(Encode(text, shift))
		case 2:
			fmt.Print("Please enter the text to decode: ")
			var text string
			_, _ = fmt.Scan(&text)

			fmt.Print("Please enter the shift: ")
			var shift rune
			_, _ = fmt.Scanln(&shift)

			fmt.Println(Decode(text, shift))
		case 3:
			return
		default:
			fmt.Println("Please choose an option 1/2/3")
		}
	}
}
