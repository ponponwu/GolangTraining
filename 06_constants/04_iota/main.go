package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
