package piscine

import (
	"github.com/01-edu/z01"
)

func main() {
	hello()
}

func hello() {
	s := "Hello World!"
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
