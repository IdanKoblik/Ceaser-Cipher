package main

import (
	"fmt"
)

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

			fmt.Println(encode(text, shift))
		} else if option == 2 {
			fmt.Print("Please enter the text to decode: ")
			var text string
			_, _ = fmt.Scan(&text)

			fmt.Print("Please enter the shift: ")
			var shift rune
			_, _ = fmt.Scanln(&shift)

			fmt.Println(decode(text, shift))
		} else if option == 3 {
			return
		} else {
			fmt.Println("Invalid option")
			return
		}
	}
}
