package main

import (
	"github.com/01-edu/z01"
)

func main() {
	LastRune("Hello!")
	LastRune("Salut!")
	LastRune("Ola!")
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	dern_lettre := []rune(s)
	return dern_lettre[len(dern_lettre)-1]
}
