package main

import "github.com/01-edu/z01"

func main() {
	printrevcomb()
}

func printrevcomb() {
	for a := 9; a >= 2; a-- {
		for b := a - 1; b >= 1; b-- {
			for c := b - 1; c >= 0; c-- {
				if a > b && b > c {
					z01.PrintRune(rune(a + 48))
					z01.PrintRune(rune(b + 48))
					z01.PrintRune(rune(c + 48))
					z01.PrintRune(rune(' '))
				}
			}

		}

	}
}
