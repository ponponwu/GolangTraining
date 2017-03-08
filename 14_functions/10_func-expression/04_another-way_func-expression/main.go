package main

import "fmt"

func main() {
	greet := makeGreeter()
	fmt.Printf("%T\n", greet)
	fmt.Printf("%T\n", greet())
	fmt.Println(greet())
}

func makeGreeter() func() string {
	return func() string {
		return "Hello world"
	}
}
