package main

import (
	"fmt"

	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		defer wg.Done()
		go func(i int) {
			fmt.Printf("i:%d", i)
		}(i)
	}

	wg.Wait()

	fmt.Println("exit")
}
