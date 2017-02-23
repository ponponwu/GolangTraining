package main

import "fmt"

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	fmt.Printf("%p, %T\n", p, p)
	fmt.Println(p)
}

func main() {
	var err error
	p, err = foo()
	fmt.Printf("%p, %T\n", p, p)
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
