package main

import "fmt"

// func main() {
// 	world()
// 	hello()
// }

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	defer fmt.Println("shit")
	hello()
	defer fmt.Println("fuck")
}
