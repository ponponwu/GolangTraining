package main

import "fmt"

func zero(z *int) {
	*z = 0
	fmt.Printf("%p\n", z)
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	zero(&x)
	fmt.Println(x)
}
