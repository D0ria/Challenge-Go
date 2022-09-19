package piscine

func IterativePower(nb int, power int) int {
	puissance := 1
	if nb == 0 {
		return 0
	}
	if power < 0 {
		return 0
	}
	if (nb < 0 && power > 0) || (nb > 0 && power < 0) {
		return 0
	}
	for i := 0; i < power; i++ {
		puissance *= nb
	}
	return puissance
}