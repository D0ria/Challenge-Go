package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			if a <= b {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b) + 48)
				z01.PrintRune(' ')

			}
		}
	}
}
