package piscine

func IsPrintable(s string) bool {
	booleen := true
	for _, val := range s {
		if val == '0' || val == '\n' {
			booleen = false
		} else {
			booleen = true
		}
	}
	return booleen
}
