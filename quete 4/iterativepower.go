package piscine

func IterativePower(nb int, power int) int {
	puissance := 1
	if nb < 0 && power < 0 {
		return 0
	} else if nb == 0 && power == 0 {
		return 1
	} else if nb == 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		for i := 0; i < power; i++ {
			puissance *= nb
		}
	}
	return puissance
}
