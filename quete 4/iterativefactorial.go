package piscine

func IterativeFactorial(nb int) int {
	if nb <= 25 {
		res := 1
		if nb == 0 {
			return 1
		} else if nb < 1 {
			return 0
		}
		for i := 1; i <= nb; i++ {
			res *= i
		}
		return res
	}
	return 0
}
