package piscine

func Any(f func(string) bool, a []string) bool {
	var booleen bool
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			booleen = true
		}
	}
	return booleen
}
