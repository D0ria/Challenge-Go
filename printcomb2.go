package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for a := 0; a <= 7; a++ {
		for b := a + 1; b <= 9; b++ {
			if a <= b {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b) + 48)
				z01.PrintRune(' ')

			}
		}
	}
}
