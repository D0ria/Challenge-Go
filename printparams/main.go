package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	michel := os.Args
	for i, j := range michel {
		if i != 0 {
			for _, k := range j {
				z01.PrintRune(k)
			}
			z01.PrintRune('\n')
		}
	}
}
