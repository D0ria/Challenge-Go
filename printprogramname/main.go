package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for michel, v := range os.Args[0] {
		if michel > 1 {
			z01.PrintRune(rune(v))
		}
	}
}
