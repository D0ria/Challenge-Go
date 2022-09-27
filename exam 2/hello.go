package main

import "github.com/01-edu/z01"

func main() {
	hello()
}

func hello() {
	s := "Hello World !"
	for _, i := range s {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
