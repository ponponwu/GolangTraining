package main

import "fmt"

func main() {
	word := "Hello"
	letter := rune(word[1])
	fmt.Println(letter)
	fmt.Println(word[1])
}
