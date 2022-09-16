package piscine

func Capitalize(s string) string {
	var vegeta string
	for i := 0; i < len(s); i++ {
		if i != 0 && (s[i-1] < 48 || s[i-1] > 57) && (s[i-1] < 65 || s[i-1] > 90) && (s[i-1] < 97 || s[i-1] > 122) && (s[i] >= 97 && s[i] <= 122) {
			vegeta += string(rune(s[i] - 32))
		} else if i != 0 && ((s[i-1] >= 97 && s[i-1] <= 122) || (s[i-1] >= 65 && s[i-1] <= 90)) && (s[i] >= 65 && s[i] <= 90) && (s[i-1] >= 48 && s[i-1] <= 57) {
			vegeta += string((rune(s[i] + 32)))
		} else if i == 0 && (s[i] >= 97 && s[i] <= 122) {
			vegeta += string(rune(s[i] - 32))
		} else {
			vegeta += string((rune(s[i])))
		}
	}
	return vegeta
}
