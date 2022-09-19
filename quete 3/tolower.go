package piscine

func ToLower(s string) string {
	var a string
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			a += string((rune(s[i]) + 32))
		} else {
			a += string((rune(s[i])))
		}
	}
	return a
}
