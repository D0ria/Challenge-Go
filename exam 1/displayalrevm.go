package main

import (
	"github.com/01-edu/z01"
)

func main() {
	displayalphamrev()
}

func displayalphamrev() {
	for i := 122; i > 97; i-- {
		z01.PrintRune(rune(i))
		z01.PrintRune(rune(i - 33))
		i--
	}
}
