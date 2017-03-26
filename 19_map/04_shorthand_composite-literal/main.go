package main

import "fmt"

func main() {
	myGreeting := map[string]string{}
	fmt.Println(myGreeting)

	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}
