package piscine

func AppendRange(min, max int) []int {
	var s1 []int
	if min < max {
		for i := min; i < max; i++ {
			s1 = append(s1, i)
		}
	}
	return s1
}
