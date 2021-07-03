package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var textReverse string
	var updatedText string
	updatedText = strings.ToLower(text)
	updatedText = strings.ReplaceAll(updatedText, " ", "")

	for i := len(updatedText) - 1; i >= 0; i-- {
		textReverse += string(updatedText[i])
	}

	if updatedText == textReverse {
		fmt.Println(text, "Is Palindrome")
	} else {
		fmt.Println(text, "Is not Palindrome")
	}
}

func main() {
	slice := []string{"hello", "there", "general", "kenobi"}
	fmt.Println(slice)

	for i, value := range slice {
		fmt.Println(i, value)
	}

	// Return value only
	for _, value := range slice {
		fmt.Println(value)
	}

	// Return index only
	for i := range slice {
		fmt.Println(i)
	}

	// Challenge. Detect a palindrome.
	isPalindrome("ama")
	isPalindrome("amor a roma")

	isPalindrome("Ama")
	isPalindrome("Anita lava la tina")
}
