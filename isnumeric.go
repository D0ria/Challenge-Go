package piscine

func IsNumeric(s string) bool {
	bool := true
	val1 := []rune(s)
	for i := range val1 {
		if val1[i] >= 48 && val1[i] <= 57 {
		} else {
			bool = false
		}
	}
	return bool
}
