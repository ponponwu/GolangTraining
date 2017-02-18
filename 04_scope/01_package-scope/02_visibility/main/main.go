package main

import (
	"fmt"
	"github.com/ponponwu/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vi.MyName)
	vi.PrintVar()
}
