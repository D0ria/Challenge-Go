package piscine

import "github.com/01-edu/z01"

func main() {
	printdigits()
}

func printdigits() {
	for i := 48; i < 57; i++ {
		z01.PrintRune(rune(i))
	}
}
