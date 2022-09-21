package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var result string
	for i := len(args) - 1; i >= 0; i-- {
		result += args[i] + "\n"
	}
	for _, j := range result {
		z01.PrintRune(j)
	}
}
