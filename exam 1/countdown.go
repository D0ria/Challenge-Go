package piscine

import "github.com/01-edu/z01"

func main() {
	countdown()
}

func countdown() {
	for i := 57; i > 48; i-- {
		z01.PrintRune(rune(i))
	}
}
