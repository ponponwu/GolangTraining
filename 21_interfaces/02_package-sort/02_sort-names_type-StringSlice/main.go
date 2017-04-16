package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	// fmt.Println(s)
	// fmt.Printf("%T", s)
	sort.StringSlice(s).Sort()
	// sort.Sort(sort.StringSlice(s))
	// fmt.Println(s)
	s = sort.StringSlice(s)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	t := sort.StringSlice(s)
	fmt.Printf("%T\n", t)
	fmt.Printf("%T\n", sort.StringSlice(s))

	fmt.Printf("s reversed: %T\n", sort.Reverse(sort.StringSlice(s)))
	i := sort.Reverse(sort.StringSlice(s))
	fmt.Println(i)
	fmt.Printf("%T\n", i)

}
