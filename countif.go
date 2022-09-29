package piscine

func CountIf(f func(string) bool, tab []string) int {
	var nb int
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			nb += i
		}
	}
	return nb
}
