package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i < len(s)-len(toFind); i++ {
		for j := range toFind {
			if s[i+j] != toFind[j] {
				break
			}
			if j == len(toFind)-1 {
				return i
			}
		}
	}
	return -1
}
