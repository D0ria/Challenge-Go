package piscine

import (
	"github.com/01-edu/z01"
)

func main() {
	printreversealphabet()
}

func printreversealphabet() {
	for i := 122; i > 97; i-- {
		z01.PrintRune(rune(i))
	}
}
