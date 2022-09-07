package main

import "github.com/01-edu/z01"

func main() {
	for i := 57; i >= 48; i-- {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
