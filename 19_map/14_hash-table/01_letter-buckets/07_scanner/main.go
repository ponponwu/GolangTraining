package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "It is not the critic who counts; not the man who points out how"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
