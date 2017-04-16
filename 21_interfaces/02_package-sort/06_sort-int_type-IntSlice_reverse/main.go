package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	// n = sort.IntSlice(n)
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	// i := sort.Reverse(sort.IntSlice(n))
	// fmt.Println(n)
	// fmt.Printf("%T\n", i)
	// sort.Sort(i)
	// fmt.Println(i)
	// sort.Sort(i)
	// fmt.Println(n)
}
