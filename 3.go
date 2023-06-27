package main

import "fmt"

func reverseString(s *string) {
	runes := []rune(*s)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*s = string(runes)
}

func main() {
	str := "Hello, World!"
	fmt.Println("Original:", str)

	reverseString(&str)
	fmt.Println("Reversed:", str)
}
