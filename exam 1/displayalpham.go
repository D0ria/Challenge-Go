package main

import "github.com/01-edu/z01"

func main() {
	Displayalpham()
}

func Displayalpham() {
	for i := 97; i < 122; i++ {
		z01.PrintRune(rune(i))
		z01.PrintRune(rune(i - 31))
		i++
	}
}
