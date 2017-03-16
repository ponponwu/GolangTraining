package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	fmt.Println(greeting[3])
	// greeting[3] = "Suprabadham"

	fmt.Println("Len:", len(greeting), "Capacity:", cap(greeting), "Value: ", greeting[0])
	// fmt.Println(greeting[2])
}
