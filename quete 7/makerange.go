package piscine

func MakeRange(min, max int) []int {
	var s1 []int
	cpt := 0
	if min >= max {
		return nil
	} else {
		s1 = make([]int, max-min)
		for i := min; i < max; i++ {
			s1[cpt] = i
			cpt += 1
		}
	}
	return s1
}
