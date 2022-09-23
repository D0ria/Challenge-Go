package piscine

func SplitWhiteSpaces(s string) []string {
	var s1 []string
	var mot string
	for j, i := range s {
		if string(i) == " " {
			s1 = append(s1, mot)
			mot = ""
		} else if j == len(s)-1 {
			s1 = append(s1, mot)
		} else {
			mot += string(i)
		}
	}
	return s1
}
