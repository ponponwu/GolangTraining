package main

import "fmt"

func fuck() {
	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

func main() {
	fuck()
}
