package piscine

func StrRev(s string) string {
	var res string
	ch := []rune(s)
	for i := len(ch) - 1; i >= 0; i-- {
		res += string(ch[i])
	}
	return res
}
