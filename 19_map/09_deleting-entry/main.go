package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"zero":         "Good morning!",
		"one":          "Bonjour!",
		"two":          "Buenos dias!",
		"morningthree": "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println(myGreeting)
}
