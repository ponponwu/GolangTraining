package main

import "fmt"

func main() {
	var name interface{} = "Sydney"
	// var name = "Sy"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
