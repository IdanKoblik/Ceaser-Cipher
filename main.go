package main

import (
	"cipher/util"
	"fmt"
)

func main() {
	for {
		fmt.Println("Welcome to Caesar Cipher decoder/encoder/cracker in Golang")

		fmt.Println("Please choose an option:")
		fmt.Println("1. Encode")
		fmt.Println("2. Decode")
		fmt.Println("3. Crack")
		fmt.Println("4. Exit")

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

			fmt.Println(util.Encode(text, shift))
		case 2:
			fmt.Print("Please enter the text to decode: ")
			var text string
			_, _ = fmt.Scan(&text)

			fmt.Print("Please enter the shift: ")
			var shift rune
			_, _ = fmt.Scan(&shift)

			fmt.Println(util.Decode(text, shift))
		case 3:
			fmt.Print("Please enter text to crack: ")
			var text string
			_, _ = fmt.Scan(&text)

			fmt.Println(util.Crack(text))
		case 4:
			return
		default:
			fmt.Println("Please choose an option 1/2/3/4")
		}
	}
}
