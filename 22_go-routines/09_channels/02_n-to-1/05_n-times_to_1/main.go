package main

import (
	"fmt"
)

func main() {

	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			fmt.Println("-----")
			fmt.Printf("%d, Done", i)
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		fmt.Println("close")
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
