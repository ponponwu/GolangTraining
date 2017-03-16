package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Good"
	greeting[1] = "Bonjour"
	greeting[2] = "bbdsf"
	greeting = append(greeting, "dfsdfsdf")
	greeting = append(greeting, "ADADSASD")
	greeting = append(greeting, "asdasdsadasd")
	greeting = append(greeting, "ffff")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
