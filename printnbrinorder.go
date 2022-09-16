package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(nb int) {
	var tabint []int
	for i := nb; i > 0; i /= 10 {
		tabint = append(tabint, i%10)
	}
	tabint = sortTab(tabint)
	for _, nb := range tabint {
		z01.PrintRune(rune(nb + 48))
	}
}

func sortTab(tab []int) []int {
	for i := 0; i < len(tab); i++ {
		for l := i; l < len(tab); l++ {
			if tab[i] < tab[l] {
				tab[i], tab[l] = tab[l], tab[i]
			}
		}
	}
	return tab
}
