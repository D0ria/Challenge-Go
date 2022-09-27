package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintComb()
}

func PrintComb() {
	for a := 0; a <= 7; a++ {
		for b := a + 1; b <= 8; b++ {
			for c := b + 1; c <= 9; c++ {
				if a < b && b < c {
					z01.PrintRune(rune(a + 48))
					z01.PrintRune(rune(b + 48))
					z01.PrintRune(rune(c + 48))
					z01.PrintRune(' ')
				}

			}

		}

	}

}
