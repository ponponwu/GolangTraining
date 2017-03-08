package main

import "fmt"

func main() {
	data := []float64{43, 56, 45, 76, 67, 12}
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
