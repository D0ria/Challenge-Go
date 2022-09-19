package piscine

func TrimAtoi(s string) int {
	w := []rune(s)
	var chiffre int
	var negatif bool
	var calcul int
	for i := 0; i < len(s); i++ {
		if w[i] >= 48 && w[i] <= 57 {
			i = len(s)
		} else if w[i] == 45 {
			negatif = true
		}
	}
	for i := 0; i < len(s); i++ {
		if w[i] >= 48 && w[i] <= 57 {
			chiffre *= 10
			chiffre += (int(w[i] - 48))
		}
	}
	if negatif == true {
		calcul = chiffre * 2
		chiffre -= calcul
	}
	return chiffre
}
